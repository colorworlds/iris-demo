package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"IRIS_WEB/tool/db"
)

var Conf *Config

// 初始化解析参数
var _path string

func init() {
	flag.StringVar(&_path, "c", "./config.yaml", "default config path")
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

// 总的配置
type Config struct {
	Server ServerConf `yaml:"server"`
	Mysql  db.MysqlConf  `yaml:"mysql"`
	Redis  db.RedisConf  `yaml:"redis"`
}

// 服务的配置
type ServerConf struct {
	Port int `yaml:"port"`
}
