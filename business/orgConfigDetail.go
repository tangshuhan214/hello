package business

import (
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strconv"
)

func GetOrgInfoConfigDetail(configId int) []models.POrgConfigDetail {
	qb, _ := orm.NewQueryBuilder("mysql")
	str := "1=1"
	if configId != 0 {
		str += " AND config_id = " + strconv.Itoa(configId)
	}
	str += " AND status != '1'"

	//拼装SQL，带分页查询。接收JSON内的String需要先转成float64再传成int传入拼装方法
	qb.Select("*").From("p_org_config_detail").Where(str)
	var list []models.POrgConfigDetail
	_, _ = orm.NewOrm().Raw(qb.String()).QueryRows(&list)
	return list
}
