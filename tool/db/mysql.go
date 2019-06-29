package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var MysqlDB *gorm.DB

type MysqlConf struct {
	Dsn     string `yaml:"dsn"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

// 初始化mysql
func InitMysql(conf *MysqlConf) (err error){
	MysqlDB, err = gorm.Open("mysql", conf.Dsn)

	if err == nil {
		MysqlDB.DB().SetMaxIdleConns(conf.MaxIdle)
		MysqlDB.DB().SetMaxOpenConns(conf.MaxOpen)
		MysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)

		err = MysqlDB.DB().Ping()
	}
	return
}

// 关闭mysql
func CloseMysql() {
	if MysqlDB != nil {
		MysqlDB.Close()
	}
}
