package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"hello/models"
	_ "hello/routers"
)

func main() {
	beego.Run()

	//开启跨域访问
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}

//初始化
func init() {
	//注册数据驱动
	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql"))
	//注册模型
	orm.RegisterModel(new(models.User))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	_ = orm.RunSyncdb("default", false, false)
}

