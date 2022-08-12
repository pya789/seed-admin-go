package entity

import "time"

type AdminUploads struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	Name        string    `json:"name" xorm:"comment('文件名')"`
	Url         string    `json:"url" xorm:"comment('地址')"`
	Type        int       `json:"type" xorm:"comment('分类')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
