// usage: redispool.Get() can get a redis connection

package redispool

import (
  "github.com/garyburd/redigo/redis"
  "github.com/astaxie/beego"
  "time"
)

type RedisPool struct {
  redis.Pool
}

var (
  RedisClient *RedisPool
)

func init(){
  RedisClient = &RedisPool{
    redis.Pool{
      MaxIdle: 8,
      IdleTimeout: 240 * time.Second,
      Dial: func () (redis.Conn, error) {
          c, err := redis.Dial("tcp", beego.AppConfig.String("redis::url"))
          if err != nil {
              return nil, err
          }
          c.Do("SELECT", beego.AppConfig.String("redis::db"))
          return c, err
      },
      TestOnBorrow: func(c redis.Conn, t time.Time) error {
          _, err := c.Do("PING")
          return err
      },
    },
  }
}

func (this *RedisPool) Get() redis.Conn {
  return this.Pool.Get()
}

func Get() redis.Conn {
  return RedisClient.Get()
}
