package db

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Locker struct {
	Key    string
	Error  error
}

func Lock(key string) (locker *Locker) {
	locker = &Locker{Key: key}

	reply, _ := redis.String(OnceRedis("SET", key, 1, "EX", 60, "NX"))

	if reply != "OK" {
		locker.Error = errors.New("locker failed.")
	}
	return
}

func TryLock(key string, timeout time.Duration) (locker *Locker) {
	locker = &Locker{Key: key}

	start := time.Now()
	for time.Now().Sub(start) < timeout {
		reply, _ := redis.String(OnceRedis("SET", key, 1, "EX", 60, "NX"))

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
		OnceRedis("DEL", lock.Key)
	}
}
