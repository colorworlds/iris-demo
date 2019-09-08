package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var mysqlDB *gorm.DB

// 初始化mysql
func StartMysql(dsn string, maxIdle, maxOpen int) (err error){
	mysqlDB, err = gorm.Open("mysql", dsn)

	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(maxIdle)
		mysqlDB.DB().SetMaxOpenConns(maxOpen)
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	}

	return
}

// 获取mysql连接
func GetMysql() *gorm.DB {
	return mysqlDB
}

// 关闭mysql
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
