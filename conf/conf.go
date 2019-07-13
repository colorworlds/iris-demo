package conf

import (
	"IRIS_WEB/utility/db"
	"IRIS_WEB/utility/log"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

// 总的配置
type Config struct {
	Server ServerConf     `yaml:"server"`
	Logger log.LoggerConf `yaml:"logger"`
	Mysql  db.MysqlConf   `yaml:"mysql"`
	Redis  db.RedisConf   `yaml:"redis"`
}

// 服务的配置
type ServerConf struct {
	Port int `yaml:"port"`
}
