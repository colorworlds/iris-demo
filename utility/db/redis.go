package db

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisPool *redis.Pool

type RedisConf struct {
	Addr    string `yaml:"addr"`
	DB      int    `yaml:"db"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

// 初始化redis
func InitRedis(conf *RedisConf) (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxOpen,
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Addr, redis.DialDatabase(conf.DB))
		},
	}

	conn := GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed.")
	}

	return
}

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// 关闭redis
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}

