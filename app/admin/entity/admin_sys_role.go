package entity

import "time"

type AdminSysRole struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"` // 主键
	Name        string    `json:"name" xorm:"notnull comment('角色名称')"`
	Label       string    `json:"label" xorm:"comment('角色标签')"`
	Remark      string    `json:"remark" xorm:"comment('备注')"`
	Relevance   int       `json:"relevance" xorm:"bool notnull default(1) comment('上下级数据权限是否关联0.是 1.否')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
