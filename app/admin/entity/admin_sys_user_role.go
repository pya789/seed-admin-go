package entity

type AdminSysUserRole struct {
	Id     int `xorm:"pk autoincr comment('主键')"`
	UserId int `xorm:"notnull comment('用户ID')"`
	RoleId int `xorm:"notnull comment('角色ID')"`
}
