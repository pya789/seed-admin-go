package entity

import "time"

type AdminSysDictType struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	Name        string    `json:"name" xorm:"comment('字典名称')"`
	Type        string    `json:"type" xorm:"comment('字典类型')"`
	Status      int       `json:"status" xorm:"bool notnull default(0) comment('状态0.正常 1.禁用')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
