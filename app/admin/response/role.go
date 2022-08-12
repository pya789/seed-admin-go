package response

type Role struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Label     string `json:"label"`
	Remark    string `json:"remark"`
	Relevance int    `json:"relevance"`
	Status    int    `json:"status"`
	MenuIds   []int  `json:"menuIds"`
	DeptIds   []int  `json:"deptIds"`
}
