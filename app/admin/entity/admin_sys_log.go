package entity

import "time"

type AdminSysLog struct {
	Id          int       `json:"id" xorm:"pk autoincr comment('主键')"`
	UserId      int       `json:"userId" xorm:"notnull comment('角色ID')"`
	Method      string    `json:"method" xorm:"notnull comment('请求方式')"`
	Action      string    `json:"action" xorm:"notnull comment('行为')"`
	Ip          string    `json:"ip" xorm:"comment('IP')"`
	StatusCode  int       `json:"statusCode" xorm:"comment('响应状态')"`
	Params      string    `json:"params" xorm:"longtext comment('参数')"`
	Results     string    `json:"results" xorm:"longtext comment('响应结果')"`
	CreatedTime time.Time `json:"createdTime" xorm:"created comment('创建时间')"`
	UpdatedTime time.Time `json:"updatedTime" xorm:"updated comment('更新时间')"`
}
