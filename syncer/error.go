// Copyright 2017 PingCAP, Inc.
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

package syncer

import (
	"database/sql/driver"
	"net"

	"github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	tddl "github.com/pingcap/tidb/ddl"
	"github.com/pingcap/tidb/infoschema"
	tmysql "github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/terror"
	gmysql "github.com/siddontang/go-mysql/mysql"
)

func ignoreDDLError(err error) bool {
	err = originError(err)
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}

	errCode := terror.ErrCode(mysqlErr.Number)
	switch errCode {
	case infoschema.ErrDatabaseExists.Code(), infoschema.ErrDatabaseNotExists.Code(), infoschema.ErrDatabaseDropExists.Code(),
		infoschema.ErrTableExists.Code(), infoschema.ErrTableNotExists.Code(), infoschema.ErrTableDropExists.Code(),
		infoschema.ErrColumnExists.Code(), infoschema.ErrColumnNotExists.Code(),
		infoschema.ErrIndexExists.Code(), tddl.ErrCantDropFieldOrKey.Code():
		return true
	case tmysql.ErrDupKeyName:
		return true
	default:
		return false
	}
}

func isRetryableError(err error) bool {
	err = originError(err)
	if err == driver.ErrBadConn || err == gmysql.ErrBadConn {
		return true
	}

	if nerr, ok := err.(net.Error); ok {
		if nerr.Timeout() {
			return true
		}
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		switch mysqlErr.Number {
		// ER_LOCK_DEADLOCK can retry to commit while meet deadlock
		case tmysql.ErrUnknown, gmysql.ER_LOCK_DEADLOCK, tmysql.ErrPDServerTimeout, tmysql.ErrTiKVServerTimeout, tmysql.ErrTiKVServerBusy, tmysql.ErrResolveLockTimeout, tmysql.ErrRegionUnavailable:
			return true
		default:
			return false
		}
	}

	return true
}

func isBinlogPurgedError(err error) bool {
	return isMysqlError(err, tmysql.ErrMasterFatalErrorReadingBinlog)
}

func isNoSuchThreadError(err error) bool {
	return isMysqlError(err, tmysql.ErrNoSuchThread)
}

func isMysqlError(err error, code uint16) bool {
	err = originError(err)
	mysqlErr, ok := err.(*mysql.MySQLError)
	return ok && mysqlErr.Number == code
}

// originError return original error
func originError(err error) error {
	for {
		e := errors.Cause(err)
		if e == err {
			break
		}
		err = e
	}
	return err
}
