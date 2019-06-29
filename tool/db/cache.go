package db

import (
	"github.com/garyburd/redigo/redis"
	"errors"
	"time"
)

//SetCache 设置缓存
func SetCache(key, val string, ttl time.Duration) (bool, error){
	r, err := redis.String(OnceRedis("SET", key, val, "EX", ttl.Seconds()))

	if err == redis.ErrNil || err != nil {
		return false, err
	}

	if r != "OK" {
		return false, errors.New("NOT OK")
	}

	return true, nil
}

//GetCache 获取缓存
func GetCache(key string) (string, error) {
	r, err := redis.String(OnceRedis("GET", key))

	if err == redis.ErrNil || err != nil {
		return "", err
	}

	return r, nil
}