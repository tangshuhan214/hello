package business

import (
	"github.com/astaxie/beego/orm"
	"github.com/gitstliu/go-id-worker"
	"hello/models"
	"strings"
)

//查询机构列表，一次多个条目
func GetOrgInfo(params map[string]interface{}, c chan map[string]interface{}) {
	qb, _ := orm.NewQueryBuilder("mysql")

	str := "1=1"

	if params["id"] != nil {
		str += " AND id = " + params["id"].(string)
	}
	if params["name"] != nil {
		str += " AND org_name like '%" + strings.Trim(params["orgName"].(string), " ") + "%'"
	}

	//拼装SQL，带分页查询。接收JSON内的String需要先转成float64再传成int传入拼装方法
	qb.Select("*").From("p_org_info").Where(str).Limit(int(params["limit"].(float64))).Offset(int(params["start"].(float64)))
	var orgInfos []models.POrgInfo
	_, _ = orm.NewOrm().Raw(qb.String()).QueryRows(&orgInfos)
	resp := map[string]interface{}{"root": &orgInfos, "total": len(orgInfos), "status": 200}
	//time.Sleep(5 * time.Second)
	c <- resp
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
