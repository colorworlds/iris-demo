package locker

import (
	"IRIS_WEB/utility/db"
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Locker struct {
	Key    string
	Error  error
}

func Lock(key string) (locker *Locker) {
	locker = &Locker{Key: key}

	conn := db.GetRedis()
	defer conn.Close()

	r, _ := redis.String(conn.Do("SET", key, 1, "EX", 60, "NX"))

	if r != "OK" {
		locker.Error = errors.New("locker failed.")
	}

	return
}

func TryLock(key string, timeout time.Duration) (locker *Locker) {
	locker = &Locker{Key: key}

	conn := db.GetRedis()
	defer conn.Close()

	start := time.Now()
	for time.Now().Sub(start) < timeout {
		reply, _ := redis.String(conn.Do("SET", key, 1, "EX", 60, "NX"))

		if reply == "OK" {
			return
		}

		time.Sleep(time.Duration(200) * time.Millisecond)
	}

	locker.Error = errors.New("locker timeout.")
	return
}

func (lock *Locker) Close() {
	if lock.Error == nil {
		conn := db.GetRedis()
		defer conn.Close()

		conn.Do("DEL", lock.Key)
	}
}
