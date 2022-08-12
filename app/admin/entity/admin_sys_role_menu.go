package entity

type AdminSysRoleMenu struct {
	Id     int `xorm:"pk autoincr comment('主键')"`
	RoleId int `xorm:"notnull comment('角色ID')"`
	MenuId int `xorm:"notnull comment('菜单ID')"`
}
