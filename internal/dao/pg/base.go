package pg

import (
	"fmt"
	"log"
	"mall/internal/core"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var once sync.Once

func initDb() {
	// Create a file to store the logs
	logFile, err := os.OpenFile(core.GlobalConfig.Postgres.Log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}

	// Create a new logger that writes to the log file
	dbLog := logger.New(log.New(logFile, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             core.GlobalConfig.Postgres.SlowThreshold * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logger.LogLevel(core.GlobalConfig.Postgres.LogLevel),        // 日志级别
		IgnoreRecordNotFoundError: true,                                                        // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,                                                       // 禁用彩色打印
	})

	dbInst, err := gorm.Open(postgres.Open(core.GlobalConfig.Postgres.Dsn), &gorm.Config{
		Logger: dbLog,
	})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	db = dbInst
}

func GetInstance() *gorm.DB {
	if db != nil {
		return db
	}
	once.Do(func() {
		initDb()
	})
	return db

}
