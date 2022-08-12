package entity

import "time"

type AdminSysMenu struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	ParentId    int       `json:"parentId" xorm:"notnull default(0) comment('父级ID 0为根项目 主键不要设置为0')"`
	Name        string    `json:"name" xorm:"notnull comment('菜单名称')"`
	RouterName  string    `json:"routerName" xorm:"comment('路由名称')"`
	RouterPath  string    `json:"routerPath" xorm:"comment('路由地址')"`
	PagePath    string    `json:"pagePath" xorm:"comment('页面路径')"`
	Perms       string    `json:"perms" xorm:"comment('权限标识')"`
	Type        int       `json:"type" xorm:"bool notnull default(0) comment('类型0.目录 1.菜单 2.按钮')"`
	Icon        string    `json:"icon" xorm:"comment('图标')"`
	Sort        int       `json:"sort" xorm:"comment('排序')"`
	Visible     int       `json:"visible" xorm:"bool notnull default(0) comment('隐藏0.显示 1.隐藏')"`
	KeepAlive   int       `json:"keepAlive" xorm:"bool notnull default(0) comment('页面缓存0.否 1.是')"`
	Status      int       `json:"status" xorm:"bool notnull default(0) comment('状态0.正常 1.禁用')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
