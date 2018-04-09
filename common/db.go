package common

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"golangWeixin/config"
)

// DB 数据库连接
var DB *gorm.DB

func initDB() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if config.ServerConfig.Env == DevelopmentMode {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	DB = db
}

func init() {
	initDB()
}
