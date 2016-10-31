package xredis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/oikomi/FishChatServer2/common/conf"
	"golang.org/x/net/context"
	"time"
)

const (
	_module = "redis"
)

type conn struct {
	p   *Pool
	c   redis.Conn
	ctx context.Context
}

type Pool struct {
	*redis.Pool
	env string
}

func NewPool(c *conf.Redis) (p *Pool) {
	p = &Pool{env: fmt.Sprintf("[%s]%s@%s", c.Name, c.Proto, c.Addr)}
	// dt := redis.DialTimer(itime.NewTimer(c.Active))
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout))
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout))
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout))
	p.Pool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(c.Proto, c.Addr, cnop, rdop, wrop)
	}, c.Idle)
	p.IdleTimeout = time.Duration(c.IdleTimeout)
	p.MaxActive = c.Active
	return
}

func (p *Pool) Get(ctx context.Context) redis.Conn {
	return &conn{p: p, c: p.Pool.Get(), ctx: ctx}
}

func (p *Pool) Close() error {
	return p.Pool.Close()
}

func (c *conn) Err() error {
	return c.c.Err()
}

func (c *conn) Close() error {
	return c.c.Close()
}

func (c *conn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.c.Do(commandName, args...)
}

// NOTE not goroutine safe
func (c *conn) Send(commandName string, args ...interface{}) (err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.c.Send(commandName, args...)
}

func (c *conn) Flush() error {
	return c.c.Flush()
}

// NOTE not goroutine safe
func (c *conn) Receive() (reply interface{}, err error) {
	if err = c.Err(); err != nil {
		return
	}
	return c.c.Receive()
}