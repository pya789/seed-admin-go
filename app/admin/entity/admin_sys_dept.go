package entity

import "time"

type AdminSysDept struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	Name        string    `json:"name" xorm:"notnull comment('部门名称')"`
	ParentId    int       `json:"parentId" xorm:"comment('父级ID')"`
	Sort        int       `json:"sort" xorm:"comment('排序')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
