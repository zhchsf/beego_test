package controllers

import(
  "github.com/astaxie/beego"
  "beego/web/lib/mongo"
  // "gopkg.in/mgo.v2"
  "fmt"
  "gopkg.in/mgo.v2/bson"
  "beego/web/models"
)

type MongodbController struct {
  beego.Controller
}

func (c *MongodbController) URLMapping() {
  c.Mapping("Get", c.Get)
}

// @router /mongodb [get]
func (this *MongodbController) Get() {
  session := mongopool.Get()
  fmt.Println(session.LiveServers())
  defer session.Close()

  collection := session.DB("caishuo_testing").C("persons")

  err := collection.Insert(&models.Person{"wzc9", "wzc@888.com"})
  if err != nil {
    fmt.Println(err)
  }

  result := models.Person{}
  find_err := collection.Find(bson.M{"username": "wzc9"}).One(&result)
  if find_err != nil {
    fmt.Println(find_err)
  }

  fmt.Println(collection.Find(bson.M{"username": "wzc9"}).Count())

  this.Data["json"] = map[string]string{"mongo": "yes", "username": result.Username, "email": result.Email}
  this.ServeJSON()
}
