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

func (init *OrgControllers) ActionFunc() {
	action := init.Ctx.Input.Param(":action")
	data := init.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)

	c := make(chan map[string]interface{})
	if action == "getOrgInfo" {
		go business.GetOrgInfo(params, c)
	} else if action == "InsertOrgInfo" {
		var orgInfo models.POrgInfo
		_ = json.Unmarshal(init.Ctx.Input.RequestBody, &orgInfo)
		go business.InsertOrUpdateOrgInfo(&orgInfo, c)
	}
	resp := <-c
	init.Data["json"] = resp
	init.ServeJSON()
}
