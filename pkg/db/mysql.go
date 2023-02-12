package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	logger2 "github.com/Shonminh/bilibee/pkg/logger"
)

func NewDB() *gorm.DB {
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/miner_db?charset=utf8mb4&parseTime=True&loc=Local", userName, password)
	newLogger := logger.New(
		&logger2.Logger{},
		logger.Config{
			SlowThreshold:             time.Second * 3, // 慢 SQL 阈值
			LogLevel:                  logger.Info,     // 日志级别
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
