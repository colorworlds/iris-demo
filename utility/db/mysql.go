package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var mysqlDB *gorm.DB

type MysqlConf struct {
	Dsn     string `yaml:"dsn"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

// 初始化mysql
func InitMysql(conf *MysqlConf) (err error){
	mysqlDB, err = gorm.Open("mysql", conf.Dsn)

	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(conf.MaxIdle)
		mysqlDB.DB().SetMaxOpenConns(conf.MaxOpen)
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)

		err = mysqlDB.DB().Ping()
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
