package redismq

import (
  //"fmt"
  //"sync"
  //"net"
  //"net/rpc"
  "time"
  "github.com/garyburd/redigo/redis"
)

type RedisMq struct {
  RedisClient *redis.Pool
  RedisHost string
  RedisDB int
}

func initRedisMq(RedisHost string, RedisDB int) *RedisMq {
  rmq := &RedisMq{
    RedisHost : RedisHost,
    RedisDB : RedisDB,
  }
  rmq.RedisClient = &redis.Pool{
      MaxIdle:     1,
  		MaxActive:   10,
  		IdleTimeout: 180 * time.Second,
  		Dial: func() (redis.Conn, error) {
  			c, err := redis.Dial("tcp", rmq.RedisHost)
  			if err != nil {
  				return nil, err
  			}
  			// 选择db
  			c.Do("SELECT", rmq.RedisDB)
  			return c, nil
  		},
    }
    return rmq
}

func RunRedisMq(RedisHost string, RedisDB int) {
  rmq := initRedisMq(RedisHost, RedisDB)
  t := time.NewTicker(60 * time.Second)
  for {
    select {
    case <-t.C:
      readUrlFromMySQL(rmq)
    }
  }
}

func readUrlFromMySQL(rmq *RedisMq) {

}
