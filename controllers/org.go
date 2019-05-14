package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/business"
)

type OrgControllers struct {
	beego.Controller
}

func (this *OrgControllers) ActionFunc() {
	action := this.Ctx.Input.Param(":action")
	data := this.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)

	c := make(chan map[string]interface{})
	if action == "getOrgInfo" {
		go business.GetOrgInfo(params, c)
	}
	resp := <-c
	this.Data["json"] = resp
	this.ServeJSON()
}
