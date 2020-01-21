package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/check", &controllers.MainController{}, "get:ConsulCheck")
	beego.Router("/?:action", &controllers.HelloControllers{}, "post:ActionFunc")
	beego.Router("/ttt", &controllers.HelloControllers{}, "post:Goto")
	beego.Router("/pay/?:action", &controllers.PayControllers{}, "post:ActionFunc")
	beego.Router("/org/?:action", &controllers.OrgControllers{}, "post:ActionFunc")
	beego.Router("/tcp/send", &controllers.TcpControllers{}, "post:TcpOnline")
	beego.Router("/tcp/one", &controllers.TcpControllers{}, "post:One")
	beego.Router("/tcp/two", &controllers.TcpControllers{}, "post:Two")
	beego.Router("/tcp/three", &controllers.TcpControllers{}, "post:Three")
}
