package entity

import "time"

type AdminSysDictData struct {
	Id          int       `json:"id"  xorm:"pk autoincr comment('主键')"`
	Pid         int       `json:"pid" xorm:"notnull comment('主键')"`
	Label       string    `json:"label" xorm:"comment('字典标签')"`
	Value       string    `json:"value" xorm:"comment('字典值')"`
	Status      int       `json:"status" xorm:"bool notnull default(0) comment('状态0.正常 1.禁用')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
