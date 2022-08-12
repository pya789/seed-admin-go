package response

type DeptTree struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	ParentId int        `json:"parentId"`
	Sort     int        `json:"sort"`
	Children []DeptTree `json:"children"`
}
