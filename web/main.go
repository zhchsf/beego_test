package main

import (
	_ "beego/web/routers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
)


func init(){
  orm.RegisterDriver("mysql", orm.DRMySQL)
  mysql_conn := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/" + beego.AppConfig.String("mysqldb")
  orm.RegisterDataBase("default", "mysql", mysql_conn, 30, 50)

}

func main() {
	beego.Run()
}


