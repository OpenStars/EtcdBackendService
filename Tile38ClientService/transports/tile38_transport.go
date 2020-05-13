package transports

import (
	"errors"
	"log"

	"github.com/gomodule/redigo/redis"
)

func newPool(maxIdle, maxActive int, address string) (p *redis.Pool, err error) {
	p = &redis.Pool{
		MaxIdle:   maxIdle,   // max number should be kept
		MaxActive: maxActive, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				log.Printf("[Tile38-RedigoClient] Dial error %#v\n", err)
			}
			return c, err
		},
	}
	if err != nil {
		p = nil
	}
	return
}

var (
	MapPool = make(map[string]*redis.Pool)
)

func GetTile38LocationClient(host, port string) (conn redis.Conn, err error) {
	log.Printf("[GetTile38LocationClient - transport] address %s %s\n", host, port)
	p, existed := MapPool[host+":"+port]
	if !existed || p == nil {
		p, _ = newPool(1000, 10000, host+":"+port)
		if p == nil {
			return nil, err
		}
		MapPool[host+":"+port] = p
	}
	if p != nil {
		return p.Get(), nil
	}
	return nil, errors.New("Can not get client")
}
