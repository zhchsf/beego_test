package controllers

import (
  "github.com/astaxie/beego"
  "fmt"
  "math/rand"
  "github.com/garyburd/redigo/redis"
  "beego/web/lib/redis"
)

type RedisController struct {
  beego.Controller
}

func (c *RedisController) URLMapping() {
  c.Mapping("Get", c.Get)
}

// @router /redis/go [get]
func (this *RedisController) Get() {
  conn := redispool.Get()
  defer conn.Close()
  // conn.Do("SET", "test", 2)
  test_value, _ := redis.String(conn.Do("GET", "test"))
  fmt.Println(test_value)
  this.Data["json"] = map[string]string{"test": test_value, "aa": "ss"}
  this.ServeJSON()
}

func RandNum() int32{
  r := rand.New(rand.NewSource(77))
  return r.Int31()
}
