package redis

import (
	"os"

	"github.com/gomodule/redigo/redis"
)


func NewConnection() *redis.Pool {

  redisPool := &redis.Pool{
    MaxIdle: 10,
    Dial: func() (redis.Conn, error) {

      return redis.Dial("tcp", os.Getenv("REDIS_URI"))
    },
  }

  return redisPool
}
