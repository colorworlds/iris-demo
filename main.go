package main

import (
	"IRIS_WEB/conf"
	"IRIS_WEB/http"
	"IRIS_WEB/utility/db"
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Parse()
	checkErr("InitConfig", conf.InitConfig())

	defer db.CloseMysql()
	checkErr("InitMysql", db.InitMysql(&conf.Conf.Mysql))

	defer db.CloseRedis()
	checkErr("InitRedis", db.InitRedis(&conf.Conf.Redis))
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf(errMsg + " Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	http.RunIris(conf.Conf.Server.Port)
}