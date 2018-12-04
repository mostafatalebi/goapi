package main

import (
	"github.com/go-redis/redis"
	"os"
	"fmt"
	"time"
)

type RedisAdapter struct{
	engine *redis.Client
	conn_err error 
	connected int // -1 not initialized, 1 ok, 0 error
}
const (
	REDIS_HOST = "127.0.0.1:6301"
)

func (rd *RedisAdapter) Connect() {
	if rd.connected == -1 {
		rd.engine = redis.NewClient(&redis.Options{
			Addr: REDIS_HOST,
			Password: "", 
			DB: 0,
		})
		_, err := rd.engine.Ping().Result()
		if err != nil {
			rd.connected = 0
			fmt.Println(err)
			os.Exit(1)
		}
	}
	rd.connected = 1
	rd.conn_err = nil
}

func (rd *RedisAdapter) Set(key string, value string, exp time.Duration) *redis.StatusCmd {
	rd.Connect()

	return rd.engine.Set(key, value, exp)
}