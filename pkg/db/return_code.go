package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	// ErrDuplicateEntryCode
	ErrDuplicateEntryCode = 1062
)

// MysqlErrCode 根据mysql错误信息返回错误代码
func MysqlErrCode(err error) int {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return 0
	}
	return int(mysqlErr.Number)
}

// IsMysqlDuplicateErr 检查当前的err是否是重复键错误
func IsMysqlDuplicateErr(err error) bool {
	return MysqlErrCode(errors.Cause(err)) == ErrDuplicateEntryCode
}
