package business

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strings"
	"time"
)

type UserBiz struct {

}

func (userBiz *UserBiz) TimeOut() {
	time.Sleep(5 * time.Second)
	fmt.Print("2222222222222")
}

func (userBiz *UserBiz)GetUserBiz(aaa map[string]interface{}, c chan map[string]interface{}) {
	qb, _ := orm.NewQueryBuilder("mysql")

	str := "1=1"

	if aaa["id"] != nil {
		str += " AND id = " + aaa["id"].(string)
	}
	if aaa["name"] != nil {
		str += " AND name like '%" + strings.Trim(aaa["name"].(string), " ") + "%'"
	}
	if aaa["pwd"] != nil {
		s := aaa["pwd"].(string)
		trim := strings.Trim(s, " ")
		if trim != "" {
			str += " AND pwd = " + trim
		}
	}

	//拼装SQL，带分页查询。接收JSON内的String需要先转成float64再传成int传入拼装方法
	qb.Select("*").From("user").Where(str).Limit(int(aaa["limit"].(float64))).Offset(int(aaa["offset"].(float64)))
	var users []models.User //用一个user集合来接收
	_, _ = orm.NewOrm().Raw(qb.String()).QueryRows(&users)

	final := Struct2Slice(users) //这里处理user结构体内万一出现一个没有定义在库表里的值

	for _, value := range final {
		value["new"] = "newField" //加入新键值对
	}

	//转成JSON发往前端数据页面，带上分页参数，带上响应状态码
	resp := map[string]interface{}{"root": &final, "total": len(users), "status": 200}
	c <- resp
}

/**
 *	这里接收一个interface{}空接口将切片结构体集合转为一个切片map集合
 */
func Struct2Slice(list interface{}) []map[string]interface{} {
	var final []map[string]interface{}
	switch list.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case []models.User:
		for _, user := range list.([]models.User) {
			jsonBytes, _ := json.Marshal(user)   //将结构体转为JSON字符串
			var only map[string]interface{}      //这里做两次转换的原因是去除首字母大写
			_ = json.Unmarshal(jsonBytes, &only) //将JSON字符串转为JSON
			final = append(final, only)          //切片不断添加新值
		}
	}
	return final
}