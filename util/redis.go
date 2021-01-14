package util

import (
	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	host,_ 	:= GetRedisIni("redis_host")
	psw,_ 	:= GetRedisIni("redis_password")
	port,_	:= GetRedisIni("redis_port")

	rdb := redis.NewClient(&redis.Options{
		Addr:     host+":"+port,
		Password: psw, // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
