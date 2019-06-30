package cache

import (
	"IRIS_WEB/utility/db"
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

//设置缓存
func Set(key, val string, ttl time.Duration) error {
	conn := db.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, val, "EX", ttl.Seconds()))

	if err != nil {
		return err
	}

	if r != "OK" {
		return errors.New("NOT OK")
	}

	return nil
}

//获取缓存
func Get(key string) (string, error) {
	conn := db.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))

	if err != nil {
		return "", err
	}

	return r, nil
}
