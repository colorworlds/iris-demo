package main

import (
	"IRIS_WEB/conf"
	"IRIS_WEB/http"
	"IRIS_WEB/utility/db"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	checkErr("InitConfig", conf.InitConfig())

	defer db.CloseMysql()
	checkErr("InitMysql", db.InitMysql(&conf.Conf.Mysql))

	defer db.CloseRedis()
	checkErr("InitRedis", db.InitRedis(&conf.Conf.Redis))

	// 设置日志
	f, err := conf.InitFileLog("iris_web")
	checkErr("InitFileLog", err)
	defer f.Close()

	http.RunIris(conf.Conf.Server.Port)
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf(errMsg+" Error: %v\n", err)
		os.Exit(1)
	}
}
