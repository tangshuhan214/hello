package business

import (
	"github.com/astaxie/beego/orm"
	"hello/common"
	"hello/models"
	"strconv"
)

func GetOrgInfoConfig(orgId int) []map[string]interface{} {
	qb, _ := orm.NewQueryBuilder("mysql")
	str := "1=1"

	if orgId != 0 {
		str += " AND org_id = " + strconv.Itoa(orgId)
	}

	//拼装SQL，带分页查询。接收JSON内的String需要先转成float64再传成int传入拼装方法
	qb.Select("*").From("p_org_config").Where(str)
	var list []models.POrgConfig
	_, _ = orm.NewOrm().Raw(qb.String()).QueryRows(&list)

	finalMap := common.StructSlice2Map(list)
	for _, oc := range finalMap {
		oc["configDetails"] = GetOrgInfoConfigDetail(oc["id"].(int))
	}

	return finalMap
}
