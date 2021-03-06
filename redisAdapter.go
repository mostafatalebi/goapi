package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

//RedisAdapter the adapter used for interacting with Redis
type RedisAdapter struct {
	engine    *redis.Client
	connErr   error
	connected int // -1 not initialized, 1 ok, 0 error
}

const (
	redisHost = "127.0.0.1:6301"
)

func newRedisAdapter() *RedisAdapter {
	rd := RedisAdapter{}
	rd.connected = 0
	rd.connErr = nil
	return &rd
}

//Connect ...
func (rd *RedisAdapter) Connect() {
	if rd.connected != 1 {
		rd.engine = redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: "",
			DB:       0,
		})
		_, err := rd.engine.Ping().Result()
		if err != nil {
			rd.connected = 0
			fmt.Println(err)
			os.Exit(1)
		}
	}
	rd.connected = 1
	rd.connErr = nil
}

//Set ...
func (rd *RedisAdapter) Set(key string, value string, exp time.Duration) *redis.StatusCmd {
	rd.Connect()

	return rd.engine.Set(key, value, exp)
}

func (rd *RedisAdapter) Get(key string) *redis.StringCmd {
	rd.Connect()

	return rd.engine.Get(key)
}

func (rd *RedisAdapter) Del(key string) *redis.IntCmd {
	rd.Connect()

	return rd.engine.Del(key)
}

//Redis following Storage interface{}
type Redis struct {
	adapter *RedisAdapter
}

func NewRedis() *Redis {
	r := Redis{}
	r.adapter = newRedisAdapter()
	return &r
}

//Save saves the data into the redis
func (r Redis) Save(key string, value string) bool {
	status := r.adapter.Set(key, value, 0)
	if status.Val() == "OK" {
		return true
	}
	return false
}

func (r Redis) SaveWithTTL(key string, value string, ttl int) bool {
	status := r.adapter.Set(key, value, time.Duration(ttl))
	if status.Val() == "OK" {
		return true
	}
	return false
}

func (r Redis) Read(key string) string {
	status := r.adapter.Get(key)
	return status.Val()
}

func (r Redis) Del(key string) int {
	status := r.adapter.Del(key)
	return int(status.Val())
}
