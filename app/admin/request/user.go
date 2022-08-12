package request

type Login struct {
	Username  string `json:"username" validate:"required|alphaNum|minLen:1|maxLen:18"`
	Password  string `json:"password" validate:"required|alphaDash|minLen:6|maxLen:18"`
	CaptchaId string `json:"captchaId" validata:"required"`
	Captcha   string `json:"captcha" validate:"required|minLen:4|maxLen:4"`
}
type User struct {
	Username string  `json:"username" validate:"required|alphaNum|minLen:1|maxLen:18"`
	NickName *string `json:"nickName" validate:"required|minLen:1|maxLen:12"`
	Password string  `json:"password" validate:"-"`
	Avatar   string  `json:"avatar" validate:"-"`
	Phone    string  `json:"phone" validate:"-"`
	Email    string  `json:"email" validate:"-"`
	Status   int     `json:"status" validate:"-"`
	DeptId   int     `json:"deptId" validate:"-"`
	RoleIds  []int   `json:"roleIds" validate:"ints"`
}
type UserList struct {
	Username string `json:"username" validate:"-"`
	NickName string `json:"nickName" validata:"-"`
	Phone    string `json:"phone" validate:"-"`
	Status   *int   `json:"status" validate:"-"`
	DeptId   []int  `json:"deptId" validate:"-"`
	PageNum  *int   `json:"pageNum" validate:"required"`
	PageSize *int   `json:"pageSize" validate:"required"`
}

type UserRoleUpdate struct {
	Id      *int  `json:"id" validate:"required"`
	RoleIds []int `json:"roleIds" validate:"ints"`
}

type UserUpdate struct {
	Id       *int    `json:"id" validate:"required"`
	Username string  `json:"username" validate:"required|alphaNum|minLen:1|maxLen:18"`
	NickName *string `json:"nickName" validate:"required|minLen:1|maxLen:12"`
	Password string  `json:"password" validate:"-"`
	Avatar   string  `json:"avatar" validate:"-"`
	Phone    string  `json:"phone" validate:"-"`
	Email    string  `json:"email" validate:"-"`
	Status   int     `json:"status" validate:"-"`
	DeptId   int     `json:"deptId" validate:"-"`
	RoleIds  []int   `json:"roleIds" validate:"ints"`
}
type UserDel struct {
	Ids []int `json:"ids" validate:"required"`
}
type UserMove struct {
	DeptId *int  `json:"deptId" validate:"required"`
	Ids    []int `json:"ids" validate:"required"`
}
type UserAvatarUpdate struct {
	Url string `json:"url"`
}
type UserBaseInfoUpdate struct {
	NickName *string `json:"nickName" validate:"required|minLen:1|maxLen:12"`
	Phone    string  `json:"phone" validate:"-"`
	Email    string  `json:"email" validate:"-"`
}
type UserPasswordUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required|alphaDash|minLen:6|maxLen:18"`
	NewPassword string `json:"newPassword" validate:"required|alphaDash|minLen:6|maxLen:18"`
}
