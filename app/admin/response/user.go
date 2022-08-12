package response

import "time"

type UserList struct {
	Id          int       `json:"id"`
	DeptId      int       `json:"deptId"`
	Username    string    `json:"username"`
	NickName    string    `json:"nickName"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Avatar      string    `json:"avatar"`
	Status      int       `json:"status"`
	IsDeleted   int       `json:"isDeleted"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
	RoleIds     string    `json:"roleIds"`
	DeptName    string    `json:"deptName"`
}
type UserInfo struct {
	Id       int    `json:"id"`
	DeptId   int    `json:"deptId"`
	Username string `json:"username"`
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	RoleIds  []int  `json:"roleIds"`
}
type UserProfile struct {
	Id          int       `json:"id"`
	DeptId      int       `json:"deptId"`
	Username    string    `json:"username"`
	NickName    string    `json:"nickName"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Avatar      string    `json:"avatar"`
	Status      int       `json:"status"`
	IsDeleted   int       `json:"isDeleted"`
	Roles       []string  `json:"roles"`
	DeptName    string    `json:"deptName" xorm:"'name'"`
	RoleNames   []string  `json:"roleNames"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
