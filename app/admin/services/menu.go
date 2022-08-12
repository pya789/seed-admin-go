package services

import (
	"fmt"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/response"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"
)

type MenuService struct {
	authCache utils.AuthCache
}

// 获取用户菜单
func (menuService *MenuService) GetMenu(userId int) ([]response.MenuTree, error) {
	roleIds, err := menuService.authCache.GetRoleIds(userId)
	if err != nil {
		return nil, err
	}
	// 利用工具库函数把权限ID数组转成字符串
	roleIdsIn := utils.SliceToInStr(roleIds)
	// 组装sql
	sql := fmt.Sprintf(`
	SELECT m.* FROM admin_sys_menu m
	JOIN admin_sys_role_menu rm 
	on m.id = rm.menu_id AND rm.role_id in (%v)
	WHERE status = 0 GROUP BY m.id ORDER BY m.sort`, roleIdsIn)
	res, err := common.DB.QueryString(sql)
	if err != nil {
		return nil, err
	}
	menus := assemblyMenu(res)
	return menus, nil
}

// 获取用户权限
func (menuService *MenuService) GetPerms(userId int) ([]string, error) {
	roleIds, err := menuService.authCache.GetRoleIds(userId)
	perms := make([]string, 0)
	if err != nil {
		return nil, err
	}
	if len(roleIds) == 0 {
		return perms, nil
	}
	// 利用工具库函数把权限ID数组转成字符串
	roleIdsIn := utils.SliceToInStr(roleIds)
	// 组装SQL
	sql := fmt.Sprintf(`
	SELECT DISTINCT m.perms FROM admin_sys_menu m
	JOIN admin_sys_role_menu rm
	on m.id = rm.menu_id AND rm.role_id in(%v)
	WHERE m.perms IS NOT NULL AND m.perms != "" AND m.status = 0
	`, roleIdsIn)
	if err := common.DB.SQL(sql).Find(&perms); err != nil {
		return nil, err
	}
	menuService.authCache.SetMenuPerms(userId, perms)
	return perms, nil
}

// 获取全部菜单
func (*MenuService) GetAllMenu(name string, status string) ([]response.MenuTree, error) {
	// 组装sql
	sql := fmt.Sprintf(`
	SELECT * FROM admin_sys_menu 
	WHERE name LIKE '%s' AND status LIKE '%s' 
	ORDER BY sort`, "%"+name+"%", "%"+status+"%")
	res, err := common.DB.QueryString(sql)
	if err != nil {
		return nil, err
	}
	menus := assemblyMenu(res)
	return menus, nil
}

// 添加菜单
func (*MenuService) AddMenu(params *request.Menu) error {
	menu := entity.AdminSysMenu{
		ParentId:   *params.ParentId,
		Name:       *params.Name,
		RouterName: params.RouterName,
		RouterPath: params.RouterPath,
		PagePath:   params.PagePath,
		Perms:      params.Perms,
		Type:       *params.Type,
		Icon:       params.Icon,
		Sort:       params.Sort,
		KeepAlive:  params.KeepAlive,
		Status:     params.Status,
	}
	if _, err := common.DB.Insert(menu); err != nil {
		return err
	}
	return nil
}

// 编辑菜单
func (menuService *MenuService) UpdateMenu(params *request.Menu) error {
	menu := entity.AdminSysMenu{
		ParentId:   *params.ParentId,
		Name:       *params.Name,
		RouterName: params.RouterName,
		RouterPath: params.RouterPath,
		PagePath:   params.PagePath,
		Perms:      params.Perms,
		Type:       *params.Type,
		Icon:       params.Icon,
		Sort:       params.Sort,
		KeepAlive:  params.KeepAlive,
		Status:     params.Status,
		Visible:    params.Visible,
	}
	if _, err := common.DB.Where("id = ?", params.Id).AllCols().Update(menu); err != nil {
		return err
	}
	if err := menuService.authCache.UpdateAllPerm(); err != nil {
		return err
	}
	return nil
}

// 删除菜单
func (menuService *MenuService) DelMenu(params *request.MenuDel) error {
	if _, err := common.DB.Table("admin_sys_menu").In("id", params.Ids).Delete(); err != nil {
		return err
	}
	if err := menuService.authCache.UpdateAllPerm(); err != nil {
		return err
	}
	return nil
}

// 获取菜单信息
func (*MenuService) GetInfo(id int) (*entity.AdminSysMenu, error) {
	menu := &entity.AdminSysMenu{
		Id: id,
	}
	if ok, err := common.DB.Get(menu); ok {
		return menu, nil
	} else {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("获取菜单信息失败")
	}
}

// 组装菜单
func assemblyMenu(res []map[string]string) []response.MenuTree {
	menus := make([]response.MenuTree, 0, len(res))
	if len(res) == 0 {
		return menus
	}
	pid, _ := strconv.Atoi(res[0]["parent_id"])
	for _, item := range res {
		parentId, _ := strconv.Atoi(item["parent_id"])
		if pid > parentId {
			pid = parentId
		}
		isRoot := false
		if parentId == 0 {
			isRoot = true
		}
		visible := false
		if item["visible"] == "0" {
			visible = true
		}
		id, _ := strconv.Atoi(item["id"])
		t, _ := strconv.Atoi(item["type"])
		sort, _ := strconv.Atoi(item["sort"])
		keepAlive, _ := strconv.ParseBool(item["keep_alive"])
		status, _ := strconv.Atoi(item["status"])
		menus = append(menus, response.MenuTree{
			Path:      item["router_path"],
			Name:      item["router_name"],
			Component: item["page_path"],
			Id:        id,
			ParentId:  parentId,
			Meta: response.Meta{
				Icon:      item["icon"],
				Sort:      sort,
				IsRoot:    isRoot,
				Title:     item["name"],
				Type:      t,
				Perms:     item["perms"],
				KeepAlive: keepAlive,
				Status:    status,
				Visible:   visible,
			},
			Children: []response.MenuTree{},
		})
	}
	m := menuTree(menus, pid)
	return m
}

// 递归获取菜单树
func menuTree(menus []response.MenuTree, pid int) []response.MenuTree {
	var nodes = make([]response.MenuTree, 0, len(menus))
	for _, item := range menus {
		if item.ParentId == pid {
			item.Children = append(item.Children, menuTree(menus, item.Id)...)
			nodes = append(nodes, item)
		}
	}
	return nodes
}
