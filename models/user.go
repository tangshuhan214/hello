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
	Id             int       `json:"id"orm:"column(id)"`                            // 组织机构id
	ErpOrgId       string    `json:"erpOrgId"orm:"column(erp_org_id)"`              // erp机构id
	OrgCode        string    `json:"orgCode"orm:"column(org_code)"`                 // 机构编码
	OrgName        string    `json:"orgName"orm:"column(org_name)"`                 // 机构名称
	OrgShortName   string    `json:"orgShortName"orm:"column(org_short_name)"`      // 机构简称
	TargetFlag     string    `json:"targetFlag"orm:"column(target_flag)"`           // 数据来源：0：同步于ERP，1：手动创建
	Status         string    `json:"status"orm:"column(status)"`                    // 数据状态：0：删除；1：可用
	CreateTime     time.Time `json:"createTime"orm:"column(create_time)"`           // 创建时间
	CreatorId      int       `json:"creatorId"orm:"column(creator_id)"`             // 创建人员id
	UpdateTime     time.Time `json:"updateTime"orm:"column(update_time)"`           // 更新时间
	UpdatorId      int       `json:"updatorId"orm:"column(updator_id)"`             // 更新人员id
	ErpOrgParentId string    `json:"erpOrgParentId"orm:"column(erp_org_parent_id)"` // 所属公司id
	ZbOrgId        string    `json:"zbOrgId"orm:"column(zb_org_id)"`                // 总部机构id
	OrgType        string    `json:"orgType"orm:"column(org_type)"`                 // 机构类型:0:公司；1：机构
	OrgAddress     string    `json:"orgAddress"orm:"column(org_address)"`           // 机构地址
	OrgUserName    string    `json:"orgUserName"orm:"column(org_user_name)"`        // 机构联系人员
	OrgUserMobile  string    `json:"orgUserMobile"orm:"column(org_user_mobile)"`    // 机构联系人员电话
	OrgProvince    string    `json:"orgProvince"orm:"column(org_province)"`         // 机构所在省
	OrgCity        string    `json:"orgCity"orm:"column(org_city)"`                 // 机构所在市
	OrgArea        string    `json:"orgArea"orm:"column(org_area)"`                 // 机构所在区/县
}

type POrgConfig struct {
	Id         int       `json:"id"orm:"column(id)"`                  // id
	OrgId      int       `json:"orgId"orm:"column(org_id)"`           // 机构id
	ErpOrgId   string    `json:"erpOrgId"orm:"column(erp_org_id)"`    // erp机构id
	CreatorId  int       `json:"creatorId"orm:"column(creator_id)"`   // 创建人id
	CreateTime time.Time `json:"createTime"orm:"column(create_time)"` // 创建时间
}

type POrgConfigDetail struct {
	Id         int       `json:"id"orm:"column(id)"`                  // id
	ValueType  string    `json:"valueType"orm:"column(value_type)"`   // 类型：
	KeyName    string    `json:"keyName"orm:"column(key_name)"`       // 名称
	KeyValue   string    `json:"keyValue"orm:"column(key_value)"`     // 值
	ConfigId   int       `json:"configId"orm:"column(config_id)"`     // 机构信息配置id
	CreatorId  int       `json:"creatorId"orm:"column(creator_id)"`   // 创建人员
	CreateTime time.Time `json:"createTime"orm:"column(create_time)"` // 创建时间
	Status     string    `json:"status"orm:"column(status)"`          // 数据状态：0：正常；1：删除
}
