package main

import (
	"IRIS_WEB/conf"
	"IRIS_WEB/http"
	"IRIS_WEB/utility/db"
	"IRIS_WEB/utility/log"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	checkErr("InitConfig", conf.InitConfig())

	defer log.CloseLogger()
	checkErr("InitLogger", log.InitLogger(&conf.Conf.Logger))

	defer db.CloseMysql()
	checkErr("InitMysql", db.InitMysql(&conf.Conf.Mysql))

	defer db.CloseRedis()
	checkErr("InitRedis", db.InitRedis(&conf.Conf.Redis))

	http.RunIris(conf.Conf.Server.Port)
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf(errMsg + " Error: %v\n", err)
		os.Exit(1)
	}
}
