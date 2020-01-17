package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/toolbox"
	eureka "github.com/xuanbo/eureka-client"
	"hello/common"
	"hello/models"
	_ "hello/routers"
)

var SafeMap = common.NewBeeMap()

func main() {
	// 链接Eureka
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://localhost:10001/eureka/",
		App:                   "MY-MICROSERVICE",
		Port:                  8000,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
	client.Start()

	//TimerTask()

	//开启跨域访问
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// 最后一个参数必须设置为false 不然无法打印数据
	beego.InsertFilter("/*", beego.FinishRouter, FilterLog, false)
	//controllers.PoolWork.Run()
	beego.Run()
}

//初始化
func init() {
	//注册数据驱动
	// orm.RegisterDriver("mysql", orm.DR_MySQL)
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql"))
	//注册模型
	/*orm.RegisterModel(new(models.User))*/
	orm.RegisterModel(new(models.POrgInfo))
	orm.RegisterModel(new(models.POrgConfig))
	orm.RegisterModel(new(models.POrgConfigDetail))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	_ = orm.RunSyncdb("default", false, false)
}

//golang的定时任务这样触发
func TimerTask() {
	tk := toolbox.NewTask("myTask", "0/3 * * * * ?", func() error { print(); return nil })
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
}

func print() {
	SafeMap.Set("a", "b")
	fmt.Print("=================================== \n")
}

// 添加日志拦截器
var FilterLog = func(ctx *context.Context) {
	addr := ctx.Request.RemoteAddr
	url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	data := ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)
	d, _ := json.Marshal(params)
	val, _ := marshalInner(ctx.Input.Data()["json"])
	divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	topDivider := "┌" + divider
	middleDivider := "├" + divider
	bottomDivider := "└" + divider
	outputStr := "\n" + topDivider + "\n│ 请求IP:" + addr + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(d) + "\n│ 返回数据:" + string(val) + bottomDivider
	logs.Info(outputStr)
}

//序列化
func marshalInner(data interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return nil, err
	}
	return bf.Bytes(), nil
}
