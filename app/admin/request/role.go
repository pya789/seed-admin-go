package request

type Role struct {
	Name      *string `json:"name" validate:"required"`
	Label     *string `json:"label" validate:"required"`
	Remark    string  `json:"remark" validate:"-"`
	Relevance int     `json:"relevance" validate:"-"`
	MenuIds   []int   `json:"menuIds" validate:"-"`
	DeptIds   []int   `json:"deptIds" validate:"-"`
}
type RoleUpdate struct {
	Id        *int    `json:"id" validate:"required"`
	Name      *string `json:"name" validate:"required"`
	Label     *string `json:"label" validate:"required"`
	Remark    string  `json:"remark" validate:"-"`
	Relevance int     `json:"relevance" validate:"-"`
	MenuIds   []int   `json:"menuIds" validate:"-"`
	DeptIds   []int   `json:"deptIds" validate:"-"`
}
type RoleList struct {
	Name     string `form:"name" validate:"-"`
	PageNum  *int   `form:"pageNum" validate:"required"`
	PageSize *int   `form:"pageSize" validate:"required"`
}
type RoleDel struct {
	Ids []int `json:"ids" validate:"required"`
}
