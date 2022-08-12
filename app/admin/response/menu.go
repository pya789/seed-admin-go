package response

type Meta struct {
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	IsRoot    bool   `json:"isRoot"`
	Title     string `json:"title"`
	Type      int    `json:"type"`
	Perms     string `json:"perms"`
	KeepAlive bool   `json:"keepAlive"`
	Status    int    `json:"status"`
	Visible   bool   `json:"visible"`
}
type MenuTree struct {
	Path      string     `json:"path"`
	Name      string     `json:"name"`
	Meta      Meta       `json:"meta"`
	Component string     `json:"component"`
	Id        int        `json:"id"`
	ParentId  int        `json:"parentId"`
	Children  []MenuTree `json:"children"`
}
