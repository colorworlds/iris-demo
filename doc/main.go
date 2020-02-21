
import (
	//标准包
	"flag"  //用于命令行参数的解析
	"fmt"
	"os"

	//第三方包
	"github.com/lestrrat-go/file-rotatelogs"  //日志相关，没用过
	"github.com/sirupsen/logrus"  //日志库

	//导入自定义的包
	. "IRIS_WEB/config"
	"IRIS_WEB/utility/db"
	"IRIS_WEB/web"
)

func main() {
	// 初始化配置文件
	flag.Parse()
	fmt.Print("InitConfig...\r")
	checkErr("InitConfig", InitConfig())
	fmt.Print("InitConfig Success!!!\n")

	// 创建文件日志，按天分割，日志文件仅保留一周
	w, err := rotatelogs.New(Conf.LogPath)
	checkErr("CreateRotateLog", err)

	// 设置日志
	logrus.SetOutput(w)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("StartMysql...\r")
	checkErr("StartMysql", db.StartMysql(Conf.MysqlDsn, Conf.MysqlMaxIdle, Conf.MysqlMaxOpen))
	fmt.Print("StartMysql Success!!!\n")

	// 启动redis
	defer db.CloseRedis()
	fmt.Print("StartRedis...\r")
	checkErr("StartRedis", db.StartRedis(Conf.RedisAddr, Conf.RedisDB, Conf.RedisMaxIdle, Conf.RedisMaxOpen))
	fmt.Print("StartRedis Success!!!\n")

	// 开始运行iris框架
	fmt.Print("RunIris...\r")
	web.RunIris(Conf.ServerPort)
}

// 检查错误
func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}

