// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package tracing

import (
	"sync"
	"time"

	"github.com/pingcap/errors"
	"github.com/siddontang/go/sync2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pingcap/dm/dm/pb"
	"github.com/pingcap/dm/pkg/log"
)

var (
	uploadInterval = 1 * time.Minute
)

// Tracer is a tracing service provider
type Tracer struct {
	sync.Mutex
	wg     sync.WaitGroup
	closed sync2.AtomicBool

	ctx    context.Context
	cancel context.CancelFunc

	cfg Config
	cli pb.TracerClient

	tso   *tsoGenerator
	idGen *idGenerator

	jobs       map[EventType]chan *Job
	jobsClosed sync2.AtomicBool
	rpcWg      sync.WaitGroup
}

// NewTracer creates a new Tracer
func NewTracer(cfg Config) *Tracer {
	t := Tracer{
		cfg:   cfg,
		tso:   newTsoGenerator(),
		idGen: newIDGen(),
	}
	t.jobsClosed.Set(true)
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return &t
}

// Enable returns whether tracing is enabled
func (t *Tracer) Enable() bool {
	return t.cfg.Enable
}

// Start starts tracing service
func (t *Tracer) Start() {
	conn, err := grpc.Dial(t.cfg.TracerAddr, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(3*time.Second))
	if err != nil {
		log.Errorf("[tracer] grpc dial error: %s", errors.ErrorStack(err))
		return
	}
	t.cli = pb.NewTracerClient(conn)
	t.closed.Set(false)
	t.newJobChans()

	t.wg.Add(1)
	go func() {
		ctx2, _ := context.WithCancel(t.ctx)
		t.tsoProcessor(ctx2)
		// cancel()
	}()

	for _, ch := range t.jobs {
		t.wg.Add(1)
		go func(c <-chan *Job) {
			ctx2, _ := context.WithCancel(t.ctx)
			t.jobProcessor(ctx2, c)
			// cancel()
		}(ch)
	}
}

// Stop stops tracer
func (t *Tracer) Stop() {
	t.Lock()
	defer t.Unlock()
	if t.closed.Get() {
		return
	}
	if t.cancel != nil {
		t.cancel()
	}
	t.closed.Set(true)
	t.wg.Wait()
}

// PrepareRpc calls before each rpc request to tracing server
func (t *Tracer) PrepareRpc() {
	t.rpcWg.Add(1)
}

func (t *Tracer) tsoProcessor(ctx context.Context) {
	defer t.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Minute):
			t.PrepareRpc()
			err := t.syncTS()
			if err != nil {
				log.Errorf("[tracer] sync timestamp error: %s", errors.ErrorStack(err))
			}
		}
	}
}

// Assume we send GetTSO with local timestamp t1, timestamp gap between local
// and remote tso service is tx, remote tso service get GetTSO at timestamp t2
// and send response back as soon as possible. We get RPC response at local
// timestamp t3. then we have t1 + tx + ttl = t2; t2 - tx + ttl = t3
// so the real sending RPC request timestamp is t1 + tx = t2 + t1/2 - t3/2
func (t *Tracer) syncTS() error {
	defer t.rpcWg.Done()
	t.tso.Lock()
	defer t.tso.Unlock()
	t.tso.localTS = time.Now().UnixNano()
	req := &pb.GetTSORequest{Id: t.cfg.Source}
	resp, err := t.cli.GetTSO(t.ctx, req)
	if err != nil {
		return errors.Trace(err)
	}
	if !resp.Result {
		return errors.Errorf("sync ts with error: %s", resp.Msg)
	}
	currentTS := time.Now().UnixNano()
	oldSyncedTS := t.tso.syncedTS
	t.tso.syncedTS = resp.Ts + t.tso.localTS/2 - currentTS/2
	log.Debugf("[tracer] syncedTS from %d to %d, localTS %d", oldSyncedTS, t.tso.syncedTS, t.tso.localTS)
	return nil
}

func (t *Tracer) newJobChans() {
	t.closeJobChans()
	t.jobs = make(map[EventType]chan *Job)
	for i := EventSyncerBinlog; i < EventFlush; i++ {
		t.jobs[i] = make(chan *Job, t.cfg.BatchSize+1)
	}
	t.jobsClosed.Set(false)
}

func (t *Tracer) closeJobChans() {
	if t.jobsClosed.Get() {
		return
	}
	for _, ch := range t.jobs {
		close(ch)
	}
	t.jobsClosed.Set(true)
}

func (t *Tracer) jobProcessor(ctx context.Context, jobChan <-chan *Job) {
	defer t.wg.Done()

	var (
		count = t.cfg.BatchSize
		jobs  = make([]*Job, 0, count)
	)

	clearJobs := func() {
		jobs = jobs[:0]
	}

	processError := func(err error) {
		log.Errorf("[tracer] processor error: %s", errors.ErrorStack(err))
	}

	var err error
	for {
		select {
		case <-ctx.Done():
			t.PrepareRpc()
			err = t.ProcessTraceEvents(jobs)
			if err != nil {
				processError(err)
			}
			clearJobs()
			return
		case <-time.After(uploadInterval):
			t.PrepareRpc()
			err = t.ProcessTraceEvents(jobs)
			if err != nil {
				processError(err)
			}
			clearJobs()
		case job, ok := <-jobChan:
			if !ok {
				return
			}

			if job.Tp != EventFlush {
				jobs = append(jobs, job)
			}

			if len(jobs) >= count || job.Tp == EventFlush {
				t.PrepareRpc()
				err = t.ProcessTraceEvents(jobs)
				if err != nil {
					processError(err)
				}
				clearJobs()
			}
		}
	}
}

// AddJob add a job to tracer
func (t *Tracer) AddJob(job *Job) {
	if job.Tp == EventFlush {
		for i := EventSyncerBinlog; i < EventFlush; i++ {
			t.jobs[i] <- job
		}
	} else {
		t.jobs[job.Tp] <- job
	}
}

func (t *Tracer) collectBaseEvent(source, traceID string, traceType pb.TraceType) (*pb.BaseEvent, error) {
	file, line, err := GetTraceCode(3)
	if err != nil {
		return nil, errors.Trace(err)
	}

	tso := t.GetTSO()
	if traceID == "" {
		traceID = t.GetTraceID(source)
	}
	return &pb.BaseEvent{
		Filename: file,
		Line:     int32(line),
		Tso:      tso,
		TraceID:  traceID,
		Type:     traceType,
	}, nil
}