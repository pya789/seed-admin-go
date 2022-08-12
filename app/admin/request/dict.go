package request

type DictList struct {
	Name     string `form:"name" validate:"-"`
	Status   string `form:"status" validate:"-"`
	PageNum  *int   `form:"pageNum" validate:"required"`
	PageSize *int   `form:"pageSize" validate:"required"`
}
type DictAdd struct {
	Name   *string `json:"name" validate:"required"`
	Type   *string `json:"type" validate:"required"`
	Status int     `json:"status" validate:"-"`
}
type DictUpdate struct {
	Id     *int    `json:"id" validate:"required"`
	Name   *string `json:"name" validate:"required"`
	Type   *string `json:"type" validate:"required"`
	Status int     `json:"status" validate:"-"`
}
type DictDel struct {
	Ids []int `json:"ids" validate:"required"`
}
type DictInfo struct {
	Id *int `form:"id" validate:"required"`
}
type DictDataList struct {
	Pid      *int   `form:"pid" validate:"required"`
	Label    string `form:"label" validate:"-"`
	Status   string `form:"status" validate:"-"`
	PageNum  *int   `form:"pageNum" validate:"required"`
	PageSize *int   `form:"pageSize" validate:"required"`
}
type DictDataAdd struct {
	Pid    *int    `json:"pid" validate:"required"`
	Label  *string `json:"label" validate:"required"`
	Value  *string `json:"value" validate:"required"`
	Status int     `json:"status" validate:"-"`
}
type DictDataInfo struct {
	Id *int `form:"id" validate:"required"`
}
type DictDataUpdate struct {
	Id     *int    `json:"id" validate:"required"`
	Pid    *int    `json:"pid" validate:"required"`
	Label  *string `json:"label" validate:"required"`
	Value  *string `json:"value" validate:"required"`
	Status int     `json:"status" validate:"-"`
}
type DictDataDel struct {
	Ids []int `json:"ids" validate:"required"`
}
