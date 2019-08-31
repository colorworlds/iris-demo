package helper

import (
	"sync"
	"time"
)

var expireMap = make(map[string]time.Time, 100)
var expireMtx sync.RWMutex


func init() {
	go func() {
		for {
			select {
			case <-time.NewTicker(time.Minute).C:
				expireMtx.Lock()

				if len(expireMap) > 0 {
					for k, v := range expireMap {
						if time.Now().After(v) {
							delete(expireMap, k)
						}
					}
				}

				expireMtx.Unlock()
			}
		}
	}()
}

func ExpireMapAdd(key string, dur time.Duration) {
	expireMtx.Lock()
	expireMap[key] = time.Now().Add(dur)
	expireMtx.Unlock()
}

func ExpireMapKey(key string) (ok bool) {
	expireMtx.RLock()
	defer expireMtx.RUnlock()

	_, ok = expireMap[key]
	return
}

