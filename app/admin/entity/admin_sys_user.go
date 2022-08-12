package entity

import "time"

type AdminSysUser struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	DeptId      int       `json:"deptId" xorm:"comment('部门ID')"`
	Username    string    `json:"username" xorm:"comment('用户名')"`
	Password    string    `json:"password" xorm:"comment('密码')"`
	NickName    string    `json:"nickName" xorm:"default('未命名') comment('用户名称')"`
	Phone       string    `json:"phone" xorm:"comment('手机')"`
	Email       string    `json:"email" xorm:"comment('邮箱')"`
	Avatar      string    `json:"avatar" xorm:"comment('头像')"`
	Status      int       `json:"status" xorm:"bool notnull default(0) comment('状态0.正常 1.禁用')"`
	IsDeleted   int       `json:"isDeleted" xorm:"bool notnull default(0) comment('软删除0.正常 1.删除')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
