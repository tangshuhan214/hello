package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/business"
)

type PayControllers struct {
	beego.Controller
}

func (this *PayControllers) ActionFunc() {
	action := this.Ctx.Input.Param(":action")
	data := this.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)
	payInter := NewPayFactory().CreateUserFactory("alipay_pay")

	var respData map[string]interface{}
	if action == "pay" {
		respData = payInter.InsertPay(params)
	} else if action == "refund" {
		respData = payInter.RefundPay(params)
	}
	this.Data["json"] = respData
	this.ServeJSON()
}

type PayInter interface {
	InsertPay(params map[string]interface{}) map[string]interface{}
	RefundPay(params map[string]interface{}) map[string]interface{}
}

type PayFactory struct {
}

func NewPayFactory() *PayFactory {
	return &PayFactory{}
}

func (this *PayFactory) CreateUserFactory(payType string) PayInter {
	if payType == "alipay_pay" {
		return &business.ZhiFuBaoBiz{}
	} else if payType == "weixin_pay" {

	}
	return nil
}
