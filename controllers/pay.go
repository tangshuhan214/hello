package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	beeLogger "github.com/beego/bee/logger"
	"hello/business"
)

type PayControllers struct {
	beego.Controller
}

func (pay *PayControllers) ActionFunc() {
	logger := beeLogger.Log
	action := pay.Ctx.Input.Param(":action")
	data := pay.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)
	payInter := NewPayFactory().CreateUserFactory("alipay_pay")

	var respData map[string]interface{}
	switch action {
	case "pay":
		respData = payInter.InsertPay(params)
	case "refund":
		respData = payInter.RefundPay(params)
	default:
		logger.Error("未定义方法类型。")
	}
	pay.Data["json"] = respData
	pay.ServeJSON()
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
