package db

import (
	"fmt"
	"mall/internal/core"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance map[string]*gorm.DB

func initMysql() {
	mysqlConfig := core.GlobalConfig.Mysql

	tmpInstance := make(map[string]*gorm.DB, 0)
	for _, conf := range mysqlConfig {
		dsn := conf.Dsn
		db, err := gorm.Open(mysql.Open((dsn)))
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
