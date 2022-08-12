package request

type LogList struct {
	UserName string `form:"username" validate:"-"`
	Method   string `form:"method" validate:"-"`
	Action   string `form:"action" validate:"-"`
	Ip       string `form:"ip" validate:"-"`
	PageNum  *int   `form:"pageNum" validate:"required"`
	PageSize *int   `form:"pageSize" validate:"required"`
}
type LogDel struct {
	Ids []int `json:"ids" validate:"required"`
}
