package main

import (
	"IRIS_WEB/conf"
	"IRIS_WEB/http"
	"IRIS_WEB/utility/db"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	if err := conf.InitConfig(); err != nil {
		fmt.Printf("InitConfig Error: %v", err)
		return
	}

	defer db.CloseMysql()
	if err := db.InitMysql(&conf.Conf.Mysql); err != nil {
		fmt.Printf("InitMysql Error: %v", err)
		return
	}

	defer db.CloseRedis()
	if err := db.InitRedis(&conf.Conf.Redis); err != nil {
		fmt.Printf("InitRedis Error: %v", err)
		return
	}

	http.RunIris(conf.Conf.Server.Port)
}
