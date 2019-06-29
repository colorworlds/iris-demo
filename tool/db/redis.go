package db

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

var RedisPool *redis.Pool

type RedisConf struct {
	Addr    string `yaml:"addr"`
	DB      int    `yaml:"db"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

// 初始化redis
func InitRedis(conf *RedisConf) (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxOpen,
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.Addr, redis.DialDatabase(conf.DB))
		},
	}

	redisConn := RedisPool.Get()
	r, _ := redis.String(redisConn.Do("PING"))
	redisConn.Close()

	if r != "PONG" {
		err = errors.New("redis connect failed.")
	}
	return
}

// execute redis query once
func OnceRedis(commandName string, args ...interface{}) (interface{}, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return conn.Do(commandName, args ...)
}


// 关闭redis
func CloseRedis() {
	if RedisPool != nil {
		RedisPool.Close()
	}
}

