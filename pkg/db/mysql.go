package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"

	"github.com/Shonminh/bilibee/pkg/logger"
)

func NewDB() *gorm.DB {
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/miner_db?charset=utf8mb4&parseTime=True&loc=Local", userName, password)
	newLogger := logger2.New(
		&logger.Logger{},
		logger2.Config{
			SlowThreshold:             time.Second * 3, // 慢 SQL 阈值
			LogLevel:                  logger2.Info,    // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,            // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

type txFunc func(c context.Context) error

const mysqlDbInstanceKey = "mysql_db_instance_key"

var ErrDBNil = errors.New("ErrDBNil")
var ErrGetDBFailed = errors.New("ErrGetDBFailed")

func Transaction(ctx context.Context, f txFunc) error {
	db := GetDb(ctx)
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	newCtx := SetDbContext(ctx, tx)
	if err := f(newCtx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetDb(ctx context.Context) *gorm.DB {
	val := ctx.Value(mysqlDbInstanceKey)
	if val == nil {
		logger.LogPanic(ErrDBNil)
	}
	db, ok := val.(*gorm.DB)
	if !ok {
		logger.LogPanic(errors.Wrapf(ErrGetDBFailed, "val=[%+v]", val).Error())
	}
	return db
}

func SetDbContext(parent context.Context, db *gorm.DB) context.Context {
	return context.WithValue(parent, mysqlDbInstanceKey, db)
}
