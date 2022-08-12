package request

type Menu struct {
	Id         int     `json:"id" validate:"-"`
	Type       *int    `json:"type" validate:"required"`
	ParentId   *int    `json:"parentId" validate:"required"`
	Name       *string `json:"name" validate:"required"`
	RouterName string  `json:"routerName" validate:"-"`
	RouterPath string  `json:"routerPath" validate:"-"`
	PagePath   string  `json:"pagePath" validate:"-"`
	Perms      string  `json:"perms" validate:"-"`
	Icon       string  `json:"icon" validate:"-"`
	Sort       int     `json:"sort" validate:"-"`
	KeepAlive  int     `json:"keepAlive" validate:"-"`
	Status     int     `json:"status" validate:"-"`
	Visible    int     `json:"visible"  validate:"-"`
}

type MenuDel struct {
	Ids []int `json:"ids" validate:"required"`
}
