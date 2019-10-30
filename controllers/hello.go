package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hello/business"
	"hello/common"
)

// debug.SetMaxThreads(num + 1000) //设置最大线程数
// 注册工作池，传入任务
// 参数1 worker并发个数
var PoolWork = common.NewWorkerPool(100 * 100 * 20)

type HelloControllers struct {
	beego.Controller
}

type Score struct {
	Ctrl *HelloControllers
}

func (s *Score) Do() {
	data := s.Ctrl.Ctx.Input.RequestBody
	param := common.ParamsInCtrl(data)
	s.Ctrl.Data["json"] = param
	s.Ctrl.ServeJSON()
}

func (hello *HelloControllers) Goto() {
	go func() {
		sc := &Score{Ctrl: hello}
		PoolWork.JobQueue <- sc
	}()
}

func (hello *HelloControllers) ActionFunc() {
	//action := hello.Ctx.Input.Param(":action")
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
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)

	url := beego.AppConfig.String("urls")
	respData := common.PostJson(url, params)

	//GO语言的强转，一次只能转一个类型，一步一步的来
	list := respData["result"].(map[string]interface{})["list"].([]interface{})

	for _, value := range list {
		i := value.(map[string]interface{})
		fmt.Print(i["id"].(string) + "\n")
	}

	/*myChan := make(chan map[string]interface{}, 2)
	userBiz := NewUserFactory().CreateUserFactory("userBiz")

	resp := map[string]interface{}{}
	if action == "getUser" {
		go userBiz.TimeOut(myChan)
		go userBiz.GetUserBiz(params, myChan)
		resp = <- myChan
		httpResp := <- myChan
		fmt.Print(httpResp)
		close(myChan)
	}*/
	hello.Data["json"] = list
	hello.ServeJSON()

}

type UserInter interface {
	GetUserBiz(aaa map[string]interface{}, c chan map[string]interface{})
	TimeOut(c chan map[string]interface{})
}

type UserFactory struct {
}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (this *UserFactory) CreateUserFactory(userAction string) UserInter {
	if userAction == "userBiz" {
		return &business.UserBiz{}
	} else if userAction == "payBiz" {
		return &business.PayBiz{}
	}
	return nil
}
