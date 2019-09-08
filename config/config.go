package config

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	Conf        *conf        // 静态配置
	DynamicConf *dynamicConf // 动态配置

	_path       string
	_etcd       string
	_client     *clientv3.Client
	_clientOnce sync.Once
)

const SERVER_NAME = "iris_web"

// 静态配置，程序启动后无法再做更改的参数配置
type conf struct {
	ServerPort   int    `yaml:"server_port"`
	LogPath      string `yaml:"log_path"`
	MysqlDsn     string `yaml:"mysql_dsn"`
	MysqlMaxIdle int    `yaml:"mysql_max_idle"`
	MysqlMaxOpen int    `yaml:"mysql_max_open"`
	RedisAddr    string `yaml:"redis_addr"`
	RedisDB      int    `yaml:"redis_db"`
	RedisMaxIdle int    `yaml:"redis_max_idle"`
	RedisMaxOpen int    `yaml:"redis_max_open"`
}

// 动态配置，程序运行过程中，可以动态更改的参数配置
type dynamicConf struct {
	UserDefaultName string `yaml:"user_default_name"`
	UserDefaultAge  int    `yaml:"user_default_age"`
}

// 初始化解析参数
func init() {
	flag.StringVar(&_path, "c", SERVER_NAME + ".yml", "default config path")
	flag.StringVar(&_etcd, "etcd", os.Getenv("ETCD"), "default etcd address")
}

// 优先从etcd中加载配置，没有则从配置文件中加载配置
func InitConfig() error {
	var err error
	var content []byte

	if _etcd != "" {
		content, err = fetchConfig("/config/" + SERVER_NAME, watchDynamicConfig)
	} else {
		content, err = ioutil.ReadFile(_path)
	}

	if err != nil {
		return err
	}

	if len(content) == 0 {
		return errors.New("not found nothing config")
	}

	Conf = &conf{}
	if err := yaml.Unmarshal(content, Conf); err != nil {
		return err
	}

	fmt.Printf("static config => [%#v]\n", Conf)

	DynamicConf = &dynamicConf{}
	if err := yaml.Unmarshal(content, DynamicConf); err != nil {
		return err
	}

	fmt.Printf("dynamic config => [%#v]\n", DynamicConf)

	return nil
}

// 从etcd中获取配置信息
func fetchConfig(nodePath string, watchFn func(k, v string)) ([]byte, error) {
	var err error
	var result string
	var resp *clientv3.GetResponse

	_clientOnce.Do(func() {
		c := clientv3.Config{Endpoints: strings.Split(_etcd, ";"), DialTimeout: 5 * time.Second}
		_client, err = clientv3.New(c)
	})

	if err != nil {
		return []byte(""), err
	}

	if resp, err = _client.Get(context.Background(), nodePath, clientv3.WithPrefix()); err != nil {
		return []byte(""), err
	}

	if resp == nil || resp.Kvs == nil {
		return []byte(""), errors.New("no response data")
	}

	for _, kvs := range resp.Kvs {
		if kvs != nil {
			result += fmt.Sprintf("%s: %s\n", filepath.Base(string(kvs.Key)), string(kvs.Value))
		}
	}

	if watchFn != nil {
		go func() {
			rch := _client.Watch(context.Background(), nodePath, clientv3.WithPrefix())
			for wResp := range rch {
				for _, ev := range wResp.Events {
					switch ev.Type {
					case mvccpb.PUT:
						watchFn(filepath.Base(string(ev.Kv.Key)), string(ev.Kv.Value))
					case mvccpb.DELETE:
						watchFn(filepath.Base(string(ev.Kv.Key)), "")
					}
				}
			}
		}()
	}

	return []byte(result), nil
}

// 监控动态配置，并使用值拷贝进行全部替换
func watchDynamicConfig(key, val string) {
	dc := new(dynamicConf)
	*dc = *DynamicConf

	yaml.Unmarshal([]byte(key + ": " + val), dc)

	DynamicConf = dc

	fmt.Printf("Latest dynamic config => [%#v]\n", DynamicConf)
}