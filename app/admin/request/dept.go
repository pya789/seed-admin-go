package request

type DeptInfo struct {
	Id *int `form:"id" validate:"required"`
}
type DeptAdd struct {
	Name     *string `json:"name" validate:"required"`
	ParentId *int    `json:"parentId" validate:"required"`
	Sort     int     `json:"sort" validate:"-"`
}
type DeptUpdate struct {
	Id       *int    `form:"id" validate:"required"`
	Name     *string `json:"name" validate:"required"`
	ParentId *int    `json:"parentId" validate:"required"`
	Sort     int     `json:"sort" validate:"-"`
}
type DeptDel struct {
	Pid     *int  `json:"pid" validate:"required"`
	Ids     []int `json:"ids" validate:"required"`
	UserDel bool  `json:"userDel" validate:"-"`
}
