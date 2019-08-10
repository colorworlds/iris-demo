package pool

import (
	"math"
	"time"
)

type NewFuncType func() (interface{}, error)
type PingFuncType func(interface{}) error
type CloseFuncType func(interface{}) error

type Pool struct {
	New     NewFuncType   // create a new connect
	Ping    PingFuncType  // check the connect
	Close   CloseFuncType // close the connect
	Timeout time.Duration
	ch      chan *activeConn
}

type activeConn struct {
	h interface{}
	t time.Time
	p *Pool
}

func NewPool(new NewFuncType, ping PingFuncType, close CloseFuncType, maxCap int, timeout time.Duration) *Pool {
	p := &Pool{
		New:     new,
		Ping:    ping,
		Close:   close,
		Timeout: timeout,
	}

	p.ch = make(chan *activeConn, maxCap)

	// 默认初始化三分之一的连接
	minCap := int(math.Ceil(float64(maxCap) * 0.3))
	for i := 0; i < minCap; i++ {
		p.ch <- &activeConn{h: p.New(), t: time.Now(), p: p}
	}

	return p
}

func (p *Pool) Get() (interface{}, error) {
	for {
		select {
		case conn := <-p.ch:
			if conn.t.Add(p.Timeout).Before(time.Now()) {
				p.Close(conn.h)
				continue
			}
			if err := conn.p.Ping(conn.h); err != nil {
				p.Close(conn.h)
				continue
			}
			return conn.h, nil
		case <-time.After(time.Second):
			return p.New()
		}
	}
}

func (p *Pool) Put(h interface{}) error {
	select {
	case p.ch <- &activeConn{h: h, t: time.Now(), p: p}:
		return nil
	default:
		return p.Close(h)
	}
}
