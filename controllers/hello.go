package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/business"
)

type HelloControllers struct {
	beego.Controller
}

func (hello *HelloControllers) ActionFunc() {
	action := hello.Ctx.Input.Param(":action")
	//接收JSON装入一个MAP中key为string value为空接口interface{}
	/*data := hello.Ctx.Input.RequestBody
	aaa := map[string]interface{}{}
	_ = json.Unmarshal(data, &aaa)
	bbb := map[string]interface{}{}
	bbb["status"] = 200
	bbb["resp"] = &aaa
	hello.Data["json"] = bbb
	hello.ServeJSON()*/

	//接收JSON装入一个user对象中
	//var user models.User
	//err := json.Unmarshal(hello.Ctx.Input.RequestBody, &user)
	//if err != nil {
	//	fmt.Println("json.Unmarshal is err:", err.Error())
	//}
	///*jsons, _ := json.Marshal(user)
	//hello.Data["json"] = jsons
	//hello.ServeJSON()*/
	//bbb := map[string]interface{}{}
	//bbb["resp"] = user
	//hello.Data["json"] = bbb
	//hello.ServeJSON()

	data := hello.Ctx.Input.RequestBody
	aaa := map[string]interface{}{}
	_ = json.Unmarshal(data, &aaa)

	myChan := make(chan map[string]interface{}, 2)
	userBiz := NewUserFactory().CreateUserFactory("userBiz")

	resp := map[string]interface{}{}
	if action == "getUser" {
		go business.Timeout()
		go userBiz.GetUserBiz(aaa, myChan)
		resp = <- myChan
		close(myChan)
	}
	hello.Data["json"] = resp
	hello.ServeJSON()

}

type UserInter interface {
	GetUserBiz(aaa map[string]interface{}, c chan map[string]interface{})
	TimeOut()
}

type UserFactory struct {

}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (this *UserFactory)CreateUserFactory(userAction string) UserInter {
	if userAction == "userBiz" {
		return &business.UserBiz{}
	}
	return nil
}