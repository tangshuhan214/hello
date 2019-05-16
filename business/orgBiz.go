package business

import (
	"github.com/astaxie/beego/orm"
	"github.com/gitstliu/go-id-worker"
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

	finalMap := common.Struct2Map(orgInfos)
	for _, value := range finalMap {
		config := GetOrgInfoConfig(int(value["id"].(float64)))
		if config != nil {
			value["configDetails"] = config[0]["configDetails"]
			slice := splitSlice(config[0]["configDetails"])
			for _, value2 := range slice {
				if len(value2) < 5 {
					value["weixin_status"] = 0
				} else {
					for _, value3 := range value2 {
						if value3.(models.POrgConfigDetail).KeyName == "" && value3.(models.POrgConfigDetail).KeyValue == "" {
							value["weixin_status"] = 0
							break
						}
						value["weixin_status"] = 1
					}
				}
			}

		} else {
			value["weixin_status"] = 0
			value["alipay_status"] = 0
			value["pos_status"] = 0
			value["social_status"] = 0
			value["print_ip"] = 0
			value["xianjin_status"] = 0
		}
	}

	//var aa = map[string]interface{}{"aaa": "1", "bbb": "2", "ccc": "3"}
	//var bb = map[string]interface{}{"aaa": "1", "bbb": "4", "ccc": "6"}
	//var cc = map[string]interface{}{"aaa": "1", "bbb": "6", "ccc": "6"}
	//var dd = []map[string]interface{}{aa, bb, cc}
	//slice := splitSlice(dd)
	//fmt.Print(slice)

	resp := map[string]interface{}{"root": &finalMap, "total": len(orgInfos), "status": 200}
	//time.Sleep(5 * time.Second)
	c <- resp
}

//切片分组算法，根据属性进行分组，返回一个切片
func splitSlice(list interface{}) [][]interface{} {
	v := reflect.ValueOf(list) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型") //不是切片立即抛错
	}
	l := v.Len()
	ret := make([]interface{}, l) //开始将传入切片转换为[]interface{}类型
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	returnData := make([][]interface{}, 0)
	i := 0
	var j int
	for {
		if i >= len(ret) {
			break
		}
		for j = i + 1; j < len(ret) && ret[i].(models.POrgConfigDetail).ValueType == ret[j].(models.POrgConfigDetail).ValueType; j++ {
		}

		returnData = append(returnData, ret[i:j])
		i = j
	}
	return returnData
}

func InsertOrUpdateOrgInfo(orgInfo *models.POrgInfo, c chan map[string]interface{}) {
	o := orm.NewOrm()
	err := o.Begin() //开启事务控制
	if orgInfo.Id != 0 {
		o.Update(orgInfo)
	} else {
		//ID生成器19位
		currWoker := &idworker.IdWorker{}
		currWoker.InitIdWorker(100, 1)
		newId, _ := currWoker.NextId()
		id := int(newId)
		orgInfo.Id = id
		o.Insert(orgInfo) //这里用完这个指针，ID就变成了0
		orgInfo.Id = id   //再次给ID赋值，以便前端获取
	}
	resp := map[string]interface{}{"root": orgInfo} //拼装返回参数map
	if err != nil {
		o.Rollback() //事务回滚
		resp["status"] = 500
		resp["msg"] = "新增失败！"
	} else {
		o.Commit() //事务提交
		resp["status"] = 200
		resp["msg"] = "新增成功！"
	}
	c <- resp //进入管道
}
