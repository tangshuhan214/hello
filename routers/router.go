package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/?:action", &controllers.HelloControllers{}, "post:ActionFunc")
	beego.Router("/ttt", &controllers.HelloControllers{}, "post:Goto")
	beego.Router("/pay/?:action", &controllers.PayControllers{}, "post:ActionFunc")
	beego.Router("/org/?:action", &controllers.OrgControllers{}, "post:ActionFunc")
	beego.Router("/tcp/send", &controllers.TcpControllers{}, "post:TcpOnline")
}
