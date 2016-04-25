package routers

import (
	"beego/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UsersController{})
    // beego.Include(&controllers.RedisController{})
    // beego.Router("/redis/go", &controllers.RedisController{},"get:Get")
    beego.Router("/redis", &controllers.RedisController{})

    // beego.Include(&controllers.MongodbController{})
    beego.Router("/mongodb", &controllers.MongodbController{})
}
