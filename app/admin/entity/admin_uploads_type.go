package entity

import "time"

type AdminUploadsType struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	Name        string    `json:"name" xorm:"comment('分类名称')"`
	Label       string    `json:"label" xorm:"comment('分类标识')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
