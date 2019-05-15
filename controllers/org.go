package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/business"
	"hello/models"
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
	} else if action == "InsertOrgInfo" {
		var orgInfo models.POrgInfo
		json.Unmarshal(this.Ctx.Input.RequestBody, &orgInfo)
		go business.InsertOrUpdateOrgInfo(&orgInfo, c)
	}
	resp := <-c
	this.Data["json"] = resp
	this.ServeJSON()
}
