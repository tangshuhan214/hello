package business

import (
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strings"
)

func GetOrgInfo(params map[string]interface{}) map[string]interface{} {
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
	return resp
}