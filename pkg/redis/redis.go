package redis

import (
	"time"

	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/xufwind95/go-web-base/config"
)

var (
	RedisPool *redis.Pool
)

func newPool(host string, port int, password string) *redis.Pool {
	redis.Dial("tcp", ":6379", redis.DialPassword("123"))
	return &redis.Pool{
		MaxIdle:     300,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn_, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", host, port),
				redis.DialPassword(password),
			)
			if err != nil {
				panic("connect redis failed!")
			}
			return conn_, err
		},
	}
}

func InitRedisPool(conf *config.Config) {
	RedisPool = newPool(conf.Redis.Host, conf.Redis.Port, conf.Redis.Password)
}

func CloseRedisPool() {
	RedisPool.Close()
}
