package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int    `json:"id"orm:"column(id)"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	Sex  string `json:"sex"`
}

type POrgInfo struct {
	// 组织机构id
	Id int `json:"id"orm:"column(id)"`
	// erp机构id
	ErpOrgId string `json:"erpOrgId"orm:"column(erp_org_id)"`
	// 机构编码
	OrgCode string `json:"orgCode"orm:"column(org_code)"`
	// 机构名称
	OrgName string `json:"orgName"orm:"column(org_name)"`
	// 机构简称
	OrgShortName string `json:"orgShortName"orm:"column(org_short_name)"`
	// 数据来源：0：同步于ERP，1：手动创建
	TargetFlag string `json:"targetFlag"orm:"column(target_flag)"`
	// 数据状态：0：删除；1：可用
	Status string `json:"status"orm:"column(status)"`
	// 创建时间
	CreateTime time.Time `json:"createTime"orm:"column(create_time)"`
	// 创建人员id
	CreatorId int `json:"creatorId"orm:"column(creator_id)"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"orm:"column(update_time)"`
	// 更新人员id
	UpdatorId int `json:"updatorId"orm:"column(updator_id)"`
	// 所属公司id
	ErpOrgParentId string `json:"erpOrgParentId"orm:"column(erp_org_parent_id)"`
	// 总部机构id
	ZbOrgId string `json:"zbOrgId"orm:"column(zb_org_id)"`
	// 机构类型:0:公司；1：机构
	OrgType string `json:"orgType"orm:"column(org_type)"`
	// 机构地址
	OrgAddress string `json:"orgAddress"orm:"column(org_address)"`
	// 机构联系人员
	OrgUserName string `json:"orgUserName"orm:"column(org_user_name)"`
	// 机构联系人员电话
	OrgUserMobile string `json:"orgUserMobile"orm:"column(org_user_mobile)"`
	// 机构所在省
	OrgProvince string `json:"orgProvince"orm:"column(org_province)"`
	// 机构所在市
	OrgCity string `json:"orgCity"orm:"column(org_city)"`
	// 机构所在区/县
	OrgArea string `json:"orgArea"orm:"column(org_area)"`
}

type POrgConfig struct {
	// id
	Id int `json:"id"orm:"column(id)"`
	// 机构id
	OrgId int `json:"orgId"orm:"column(org_id)"`
	// erp机构id
	ErpOrgId string `json:"erpOrgId"orm:"column(erp_org_id)"`
	// 创建人id
	CreatorId int `json:"creatorId"orm:"column(creator_id)"`
	// 创建时间
	CreateTime time.Time `json:"createTime"orm:"column(create_time)"`
}

type POrgConfigDetail struct {
	// id
	Id int `json:"id"orm:"column(id)"`
	// 类型：
	ValueType string `json:"valueType"orm:"column(value_type)"`
	// 名称
	KeyName string `json:"keyName"orm:"column(key_name)"`
	// 值
	KeyValue string `json:"keyValue"orm:"column(key_value)"`
	// 机构信息配置id
	ConfigId int `json:"configId"orm:"column(config_id)"`
	// 创建人员
	CreatorId int `json:"creatorId"orm:"column(creator_id)"`
	// 创建时间
	CreateTime time.Time `json:"createTime"orm:"column(create_time)"`
	// 数据状态：0：正常；1：删除
	Status string `json:"status"orm:"column(status)"`
}
