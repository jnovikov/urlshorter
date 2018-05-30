package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"fmt"
)

type Pool struct {
	pool *pool.Pool
}

func (p *Pool) Get() (*redis.Client, error) {
	return p.pool.Get()
}

func (p *Pool) Put(conn *redis.Client) {
	p.pool.Put(conn)
}

func NewPool(host string, port string, size int) *Pool {
	p, err := pool.New("tcp", fmt.Sprintf("%s:%s", host, port), size)
	if err != nil {
		message := fmt.Sprintf("Cant connect to redis on %s:%s with error %s", host, port, err.Error())
		panic(message)
	}

	return &Pool{p}
}
