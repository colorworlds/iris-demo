package config

import (
	"IRIS_WEB/utility/db"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *conf

// 初始化解析参数
var _path string

func init() {
	flag.StringVar(&_path, "c", "config.yml", "default config path")
}

// 从配置文件中加载配置
func InitConfig() error {
	Conf = &conf{}

	content, err := ioutil.ReadFile(_path)
	if err == nil {
		err = yaml.Unmarshal(content, Conf)
	}

	if err == nil {
		// 启动准备项

		// 初始化参数
		Conf.Params = newParams()
	}

	return err
}

// 总的配置
type conf struct {
	Server serverConf   `yaml:"server"`
	Mysql  db.MysqlConf `yaml:"mysql"`
	Redis  db.RedisConf `yaml:"redis"`

	Params *params
}

// 服务的配置
type serverConf struct {
	Port int `yaml:"port"`
}
