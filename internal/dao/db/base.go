package db

import (
	"fmt"
	"mall/internal/core"
	"mall/internal/logger"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance map[string]*gorm.DB

func initMysql() {
	mysqlConfig := core.GlobalConfig.Mysql

	tmpInstance := make(map[string]*gorm.DB, 0)
	for _, conf := range mysqlConfig {
		config := logger.GormLoggerConfig{
			SlowThreshold: conf.SlowThreshold,
			TraceLog:      conf.TraceLog,
		}
		newLogger := logger.NewGormLog(config)
		gormConfig := gorm.Config{}
		gormConfig.Logger = newLogger

		dsn := conf.Dsn
		db, err := gorm.Open(mysql.Open((dsn)), &gormConfig)
		if err != nil {
			fmt.Printf("mysql connect error: %v", err)
		}
		tmpInstance[conf.Instance] = db
	}
	dbInstance = tmpInstance

}

var once sync.Once

func GetDbInstance(db string) *gorm.DB {
	if db == "" {
		db = "default"
	}

	if dbInstance != nil {
		return dbInstance[db]
	}
	once.Do(func() {
		initMysql()
	})
	return dbInstance[db]
}
