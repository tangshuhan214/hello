package business

import (
	"github.com/astaxie/beego/orm"
	idworker "github.com/gitstliu/go-id-worker"
	"hello/common"
	"hello/models"
	"reflect"
	"strconv"
	"strings"
)

//查询机构列表，一次多个条目
func GetOrgInfo(params map[string]interface{}, c chan map[string]interface{}) {
	qb, _ := orm.NewQueryBuilder("mysql")

	str := "1=1"

	if params["id"] != nil {
		str += " AND id = " + strconv.Itoa(int(params["id"].(float64)))
	}
	if params["name"] != nil {
		str += " AND org_name like '%" + strings.Trim(params["orgName"].(string), " ") + "%'"
	}

	//拼装SQL，带分页查询。接收JSON内的String需要先转成float64再传成int传入拼装方法
	qb.Select("*").From("p_org_info").Where(str).Limit(int(params["limit"].(float64))).Offset(int(params["start"].(float64)))
	var orgInfos []models.POrgInfo
	_, _ = orm.NewOrm().Raw(qb.String()).QueryRows(&orgInfos)

	finalMap := common.StructSlice2Map(orgInfos)
	for _, value := range finalMap {
		config := GetOrgInfoConfig(value["id"].(int))
		value["weixin_status"] = 0
		value["alipay_status"] = 0
		value["pos_status"] = 0
		value["social_status"] = 0
		value["print_ip"] = 0
		value["xianjin_status"] = 1

		if config != nil {
			slice := splitSlice(config[0]["configDetails"])
			value["configDetails"] = slice
			for key, configDetail := range slice {
				if key == "weixin_pay" {
					value["weixin_status"] = checkConfig(configDetail, 5)
				}

				if key == "alipay_pay" {
					value["alipay_status"] = checkConfig(configDetail, 7)
				}

				if key == "pos_pay" {
					value["pos_status"] = checkConfig(configDetail, 6)
				}

				if key == "social_pay" {
					value["social_status"] = checkConfig(configDetail, 1)
				}

				if key == "print_pay" {
					value["print_status"] = checkConfig(configDetail, 1)
				}

			}
		}
	}

	resp := map[string]interface{}{"root": &finalMap, "total": len(orgInfos), "status": 200}
	//time.Sleep(5 * time.Second)
	c <- resp
}

//切片分组算法，根据属性进行分组，返回一个切片（无法封装[捂脸]）
func splitSlice(list interface{}) map[string][]interface{} {
	v := reflect.ValueOf(list) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型") //不是切片立即抛错
	}
	l := v.Len()
	ret := make([]interface{}, l) //开始将传入切片转换为[]interface{}类型
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	returnData := make(map[string][]interface{})
	i := 0
	var j int
	for {
		if i >= len(ret) {
			break
		}
		for j = i + 1; j < len(ret) && ret[i].(models.POrgConfigDetail).ValueType == ret[j].(models.POrgConfigDetail).ValueType; j++ {
		}

		returnData[ret[i].(models.POrgConfigDetail).ValueType] = ret[i:j]
		i = j
	}
	return returnData
}

func checkConfig(list []interface{}, num int) int {
	resp := 0
	if len(list) < num {
		resp = 0
	} else {
		for _, v := range list {
			if v.(models.POrgConfigDetail).KeyName == "" && v.(models.POrgConfigDetail).KeyValue == "" {
				resp = 0
				break
			}
			resp = 1
		}
	}
	return resp
}

func InsertOrUpdateOrgInfo(orgInfo *models.POrgInfo, c chan map[string]interface{}) {
	o := orm.NewOrm()
	_ = o.Begin() //开启事务控制

	resp := map[string]interface{}{"root": orgInfo} //拼装返回参数map
	common.TryCatch{}.Try(func() {                  //try catch 做异常捕获
		if orgInfo.Id != 0 {
			_, _ = o.Update(orgInfo)
		} else {
			//ID生成器19位
			currWoker := &idworker.IdWorker{}
			_ = currWoker.InitIdWorker(100, 1)
			newId, _ := currWoker.NextId()
			id := int(newId)
			orgInfo.Id = id
			_, _ = o.Insert(orgInfo) //这里用完这个指针，ID就变成了0
			orgInfo.Id = id          //再次给ID赋值，以便前端获取
		}
		_ = o.Commit() //事务提交
		resp["status"] = 200
		resp["msg"] = "新增成功！"
	}).CatchAll(func(err error) {
		_ = o.Rollback() //事务回滚
		resp["status"] = 500
		resp["msg"] = err.Error() //输出错误信息
	}).Finally(func() {
		c <- resp //进入管道
	})

}
