package conf

import (
	"IRIS_WEB/utility/db"
	"flag"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var Conf *Config

// 初始化解析参数
var _path string

func init() {
	flag.StringVar(&_path, "c", "config.yml", "default config path")
}

// 从配置文件中加载配置
func InitConfig() error {
	Conf = &Config{}

	content, err := ioutil.ReadFile(_path)
	if err == nil {
		err = yaml.Unmarshal(content, Conf)
	}
	return err
}

// 初始化日志
func InitFileLog(serverName string) (*os.File, error) {
	logFileName := filepath.Join(Conf.Server.LogPath, serverName+"_"+time.Now().Format("20060102")+".log")
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	golog.AddOutput(f)
	return f, nil
}

// 总的配置
type Config struct {
	Server ServerConf   `yaml:"server"`
	Mysql  db.MysqlConf `yaml:"mysql"`
	Redis  db.RedisConf `yaml:"redis"`
}

// 服务的配置
type ServerConf struct {
	Port    int    `yaml:"port"`
	LogPath string `yaml:"logPath"`
}
