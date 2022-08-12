package entity

type AdminSysRoleDept struct {
	Id     int `xorm:"pk autoincr comment('主键')"`
	RoleId int `xorm:"notnull comment('角色ID')"`
	DeptId int `xorm:"notnull comment('部门ID')"`
}
