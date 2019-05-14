package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
	"hello/business"
	"hello/common"
)

var SafeMap = common.NewBeeMap()

type OrgControllers struct {
	beego.Controller
}

func (this *OrgControllers) ActionFunc() {
	action := this.Ctx.Input.Param(":action")
	data := this.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)

	var resp map[string]interface{}
	if action == "getOrgInfo" {
		resp = business.GetOrgInfo(params)
	}
	Notaaa()
	this.Data["json"] = resp
	this.ServeJSON()
}

func Notaaa() {
	go func() {
		crontab := cron.New()
		crontab.AddFunc("0/5 * * * * ?", print)
		crontab.Start()
	}()
}

func print() {
	SafeMap.Set("a", "b")
	fmt.Print("=================================== \n")
}
