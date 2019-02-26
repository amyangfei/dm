// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tracer.proto

package pb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetTSORequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetTSORequest) Reset()         { *m = GetTSORequest{} }
func (m *GetTSORequest) String() string { return proto.CompactTextString(m) }
func (*GetTSORequest) ProtoMessage()    {}
func (*GetTSORequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{0}
}
func (m *GetTSORequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTSORequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTSORequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTSORequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTSORequest.Merge(m, src)
}
func (m *GetTSORequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTSORequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTSORequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTSORequest proto.InternalMessageInfo

func (m *GetTSORequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTSOResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ts     int64  `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (m *GetTSOResponse) Reset()         { *m = GetTSOResponse{} }
func (m *GetTSOResponse) String() string { return proto.CompactTextString(m) }
func (*GetTSOResponse) ProtoMessage()    {}
func (*GetTSOResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{1}
}
func (m *GetTSOResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTSOResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTSOResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTSOResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTSOResponse.Merge(m, src)
}
func (m *GetTSOResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTSOResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTSOResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTSOResponse proto.InternalMessageInfo

func (m *GetTSOResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *GetTSOResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetTSOResponse) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type CommonUploadResponse struct {
	Result bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *CommonUploadResponse) Reset()         { *m = CommonUploadResponse{} }
func (m *CommonUploadResponse) String() string { return proto.CompactTextString(m) }
func (*CommonUploadResponse) ProtoMessage()    {}
func (*CommonUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{2}
}
func (m *CommonUploadResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonUploadResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonUploadResponse.Merge(m, src)
}
func (m *CommonUploadResponse) XXX_Size() int {
	return m.Size()
}
func (m *CommonUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonUploadResponse proto.InternalMessageInfo

func (m *CommonUploadResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *CommonUploadResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UploadSyncerBinlogEventRequest struct {
	Events []*SyncerBinlogEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *UploadSyncerBinlogEventRequest) Reset()         { *m = UploadSyncerBinlogEventRequest{} }
func (m *UploadSyncerBinlogEventRequest) String() string { return proto.CompactTextString(m) }
func (*UploadSyncerBinlogEventRequest) ProtoMessage()    {}
func (*UploadSyncerBinlogEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{3}
}
func (m *UploadSyncerBinlogEventRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadSyncerBinlogEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadSyncerBinlogEventRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadSyncerBinlogEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadSyncerBinlogEventRequest.Merge(m, src)
}
func (m *UploadSyncerBinlogEventRequest) XXX_Size() int {
	return m.Size()
}
func (m *UploadSyncerBinlogEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadSyncerBinlogEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadSyncerBinlogEventRequest proto.InternalMessageInfo

func (m *UploadSyncerBinlogEventRequest) GetEvents() []*SyncerBinlogEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type UploadSyncerJobRequest struct {
	Events []*SyncerJobEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *UploadSyncerJobRequest) Reset()         { *m = UploadSyncerJobRequest{} }
func (m *UploadSyncerJobRequest) String() string { return proto.CompactTextString(m) }
func (*UploadSyncerJobRequest) ProtoMessage()    {}
func (*UploadSyncerJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d422d7c66fbbd8f, []int{4}
}
func (m *UploadSyncerJobRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadSyncerJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadSyncerJobRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadSyncerJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadSyncerJobRequest.Merge(m, src)
}
func (m *UploadSyncerJobRequest) XXX_Size() int {
	return m.Size()
}
func (m *UploadSyncerJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadSyncerJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadSyncerJobRequest proto.InternalMessageInfo

func (m *UploadSyncerJobRequest) GetEvents() []*SyncerJobEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTSORequest)(nil), "pb.GetTSORequest")
	proto.RegisterType((*GetTSOResponse)(nil), "pb.GetTSOResponse")
	proto.RegisterType((*CommonUploadResponse)(nil), "pb.CommonUploadResponse")
	proto.RegisterType((*UploadSyncerBinlogEventRequest)(nil), "pb.UploadSyncerBinlogEventRequest")
	proto.RegisterType((*UploadSyncerJobRequest)(nil), "pb.UploadSyncerJobRequest")
}

func init() { proto.RegisterFile("tracer.proto", fileDescriptor_6d422d7c66fbbd8f) }

var fileDescriptor_6d422d7c66fbbd8f = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x29, 0x4a, 0x4c,
	0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x12, 0x86, 0x88,
	0xc4, 0x17, 0x57, 0xe6, 0xc1, 0x25, 0x94, 0xe4, 0xb9, 0x78, 0xdd, 0x53, 0x4b, 0x42, 0x82, 0xfd,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0xbc, 0xb8, 0xf8, 0x60, 0x0a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xc0, 0xaa, 0x38,
	0x82, 0xa0, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x26, 0xb0, 0x56, 0x10, 0x13,
	0x64, 0x56, 0x49, 0xb1, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x49, 0xb1, 0x92, 0x03,
	0x97, 0x88, 0x73, 0x7e, 0x6e, 0x6e, 0x7e, 0x5e, 0x68, 0x41, 0x4e, 0x7e, 0x62, 0x0a, 0xe9, 0x26,
	0x2a, 0xf9, 0x73, 0xc9, 0x41, 0xf4, 0x06, 0x83, 0x3d, 0xe1, 0x94, 0x99, 0x97, 0x93, 0x9f, 0xee,
	0x5a, 0x96, 0x9a, 0x57, 0x02, 0x73, 0xbf, 0x2e, 0x17, 0x5b, 0x2a, 0x88, 0x5f, 0x2c, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xaa, 0x57, 0x90, 0xa4, 0x87, 0xa9, 0x1a, 0xaa, 0x48, 0xc9, 0x85,
	0x4b, 0x0c, 0xd9, 0x40, 0xaf, 0xfc, 0x24, 0x98, 0x41, 0x5a, 0x68, 0x06, 0x09, 0x21, 0x0c, 0xf2,
	0xca, 0x4f, 0x42, 0x31, 0xc5, 0x68, 0x1a, 0x23, 0x17, 0x5b, 0x08, 0x38, 0x74, 0x85, 0x0c, 0xb9,
	0xd8, 0x20, 0xe1, 0x25, 0x24, 0x08, 0xd2, 0x80, 0x12, 0xb8, 0x52, 0x42, 0xc8, 0x42, 0x10, 0xcf,
	0x2b, 0x31, 0x08, 0x45, 0x72, 0x89, 0xe3, 0xf0, 0x94, 0x90, 0x12, 0x48, 0x03, 0x7e, 0x1f, 0x4b,
	0x49, 0x80, 0xd4, 0x60, 0x0b, 0x57, 0x25, 0x06, 0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xc7, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x00, 0x35,
	0x06, 0xa2, 0x28, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TracerClient is the client API for Tracer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracerClient interface {
	GetTSO(ctx context.Context, in *GetTSORequest, opts ...grpc.CallOption) (*GetTSOResponse, error)
	UploadSyncerBinlogEvent(ctx context.Context, in *UploadSyncerBinlogEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error)
}

type tracerClient struct {
	cc *grpc.ClientConn
}

func NewTracerClient(cc *grpc.ClientConn) TracerClient {
	return &tracerClient{cc}
}

func (c *tracerClient) GetTSO(ctx context.Context, in *GetTSORequest, opts ...grpc.CallOption) (*GetTSOResponse, error) {
	out := new(GetTSOResponse)
	err := c.cc.Invoke(ctx, "/pb.Tracer/GetTSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracerClient) UploadSyncerBinlogEvent(ctx context.Context, in *UploadSyncerBinlogEventRequest, opts ...grpc.CallOption) (*CommonUploadResponse, error) {
	out := new(CommonUploadResponse)
	err := c.cc.Invoke(ctx, "/pb.Tracer/UploadSyncerBinlogEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracerServer is the server API for Tracer service.
type TracerServer interface {
	GetTSO(context.Context, *GetTSORequest) (*GetTSOResponse, error)
	UploadSyncerBinlogEvent(context.Context, *UploadSyncerBinlogEventRequest) (*CommonUploadResponse, error)
}

func RegisterTracerServer(s *grpc.Server, srv TracerServer) {
	s.RegisterService(&_Tracer_serviceDesc, srv)
}

func _Tracer_GetTSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTSORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).GetTSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tracer/GetTSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).GetTSO(ctx, req.(*GetTSORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracer_UploadSyncerBinlogEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSyncerBinlogEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracerServer).UploadSyncerBinlogEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Tracer/UploadSyncerBinlogEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracerServer).UploadSyncerBinlogEvent(ctx, req.(*UploadSyncerBinlogEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Tracer",
	HandlerType: (*TracerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTSO",
			Handler:    _Tracer_GetTSO_Handler,
		},
		{
			MethodName: "UploadSyncerBinlogEvent",
			Handler:    _Tracer_UploadSyncerBinlogEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracer.proto",
}

func (m *GetTSORequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTSORequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *GetTSOResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTSOResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result {
		dAtA[i] = 0x8
		i++
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Ts != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTracer(dAtA, i, uint64(m.Ts))
	}
	return i, nil
}

func (m *CommonUploadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonUploadResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result {
		dAtA[i] = 0x8
		i++
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTracer(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func (m *UploadSyncerBinlogEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSyncerBinlogEventRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTracer(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UploadSyncerJobRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadSyncerJobRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTracer(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTracer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetTSORequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	return n
}

func (m *GetTSOResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	if m.Ts != 0 {
		n += 1 + sovTracer(uint64(m.Ts))
	}
	return n
}

func (m *CommonUploadResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovTracer(uint64(l))
	}
	return n
}

func (m *UploadSyncerBinlogEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTracer(uint64(l))
		}
	}
	return n
}

func (m *UploadSyncerJobRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTracer(uint64(l))
		}
	}
	return n
}

func sovTracer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTracer(x uint64) (n int) {
	return sovTracer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTSORequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTSORequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTSORequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTSOResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTSOResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTSOResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Result = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ts", wireType)
			}
			m.Ts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ts |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommonUploadResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonUploadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonUploadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Result = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UploadSyncerBinlogEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadSyncerBinlogEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSyncerBinlogEventRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &SyncerBinlogEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UploadSyncerJobRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadSyncerJobRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadSyncerJobRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTracer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &SyncerJobEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTracer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTracer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTracer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTracer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTracer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTracer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTracer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracer   = fmt.Errorf("proto: integer overflow")
)
