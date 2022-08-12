package services

import (
	"errors"
	"fmt"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/response"
	"seed-admin/common"
	"seed-admin/utils"
)

type RoleService struct {
	authCache utils.AuthCache
}

// 获取全部角色
func (*RoleService) GetAllRole(params *request.RoleList) ([]entity.AdminSysRole, int64, error) {
	role := new(entity.AdminSysRole)
	roles := make([]entity.AdminSysRole, 0)
	count, err := common.DB.Where("name LIKE ?", "%"+params.Name+"%").Count(role)
	if err != nil {
		return nil, 0, err
	}
	if err := common.DB.Where("name LIKE ?", "%"+params.Name+"%").
		Limit(*params.PageSize, (*params.PageNum-1)*(*params.PageSize)).
		Find(&roles); err != nil {
		return nil, 0, err
	}
	return roles, count, nil
}

// 获取角色信息
func (*RoleService) GetInfo(id int) (*response.Role, error) {
	role := &entity.AdminSysRole{
		Id: id,
	}
	menuIds := make([]int, 0)
	deptIds := make([]int, 0)
	if ok, err := common.DB.Get(role); ok {
		if err := common.DB.Table("admin_sys_role_menu").Where("role_id = ?", role.Id).Cols("menu_id").Find(&menuIds); err != nil {
			return nil, err
		}
		if err := common.DB.Table("admin_sys_role_dept").Where("role_id = ?", role.Id).Cols("dept_id").Find(&deptIds); err != nil {
			return nil, err
		}
		res := &response.Role{
			Id:        role.Id,
			Name:      role.Name,
			Label:     role.Label,
			Remark:    role.Remark,
			Relevance: role.Relevance,
			MenuIds:   menuIds,
			DeptIds:   deptIds,
		}
		return res, nil
	} else {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("获取角色信息失败")
	}
}

// 增加一个角色
func (*RoleService) AddRole(params *request.Role) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	role := &entity.AdminSysRole{
		Name:      *params.Name,
		Label:     *params.Label,
		Remark:    params.Remark,
		Relevance: params.Relevance,
	}
	// 插入角色表
	if _, err := session.Insert(role); err != nil {
		return err
	}
	if len(params.MenuIds) == 0 {
		return errors.New("至少需要选择一个菜单")
	}
	// 插入角色菜单关联表
	roleMenu := make([]entity.AdminSysRoleMenu, 0, len(params.MenuIds))
	for _, menuId := range params.MenuIds {
		roleMenu = append(roleMenu, entity.AdminSysRoleMenu{
			RoleId: role.Id,
			MenuId: menuId,
		})
	}
	if _, err := session.Insert(roleMenu); err != nil {
		return err
	}
	if len(params.DeptIds) == 0 {
		return session.Commit()
	}
	// 插入角色部门关联表
	roleDept := make([]entity.AdminSysRoleDept, 0, len(params.DeptIds))
	for _, deptId := range params.DeptIds {
		roleDept = append(roleDept, entity.AdminSysRoleDept{
			RoleId: role.Id,
			DeptId: deptId,
		})
	}
	if _, err := session.Insert(roleDept); err != nil {
		return err
	}
	return session.Commit()
}

// 更新角色
func (roleService *RoleService) UpdateRole(params *request.RoleUpdate) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	role := entity.AdminSysRole{
		Name:      *params.Name,
		Label:     *params.Label,
		Remark:    params.Remark,
		Relevance: params.Relevance,
	}
	// 更新角色表
	if _, err := session.Where("id = ?", params.Id).AllCols().Update(role); err != nil {
		return err
	}
	// 删除角色原有菜单权限
	if _, err := session.Table("admin_sys_role_menu").Where("role_id = ?", params.Id).Delete(); err != nil {
		return err
	}
	// 删除角色原有部门权限
	if _, err := session.Table("admin_sys_role_dept").Where("role_id = ?", params.Id).Delete(); err != nil {
		return err
	}
	if len(params.MenuIds) != 0 {
		// 增加新的角色权限
		roleMenu := make([]entity.AdminSysRoleMenu, 0, len(params.MenuIds))
		for _, menuId := range params.MenuIds {
			roleMenu = append(roleMenu, entity.AdminSysRoleMenu{
				RoleId: *params.Id,
				MenuId: menuId,
			})
		}
		if _, err := session.Insert(roleMenu); err != nil {
			return err
		}
	}
	if len(params.DeptIds) != 0 {
		// 增加新的角色部门
		roleDept := make([]entity.AdminSysRoleDept, 0, len(params.DeptIds))
		for _, deptId := range params.DeptIds {
			roleDept = append(roleDept, entity.AdminSysRoleDept{
				RoleId: *params.Id,
				DeptId: deptId,
			})
		}
		if _, err := session.Insert(roleDept); err != nil {
			return err
		}
	}
	if err := session.Commit(); err != nil {
		return err
	}
	if err := roleService.authCache.UpdateAllRole(); err != nil {
		return err
	}
	if err := roleService.authCache.UpdateAllPerm(); err != nil {
		return err
	}
	return nil
}

// 删除角色
func (roleService *RoleService) DelRole(params *request.RoleDel) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// 删除角色所属菜单权限
	if _, err := session.Table("admin_sys_role_menu").In("role_id", params.Ids).Delete(); err != nil {
		return err
	}
	// 删除角色原有部门权限
	if _, err := session.Table("admin_sys_role_dept").In("role_id", params.Ids).Delete(); err != nil {
		return err
	}
	// 删除用户此角色
	if _, err := session.Table("admin_sys_user_role").In("role_id", params.Ids).Delete(); err != nil {
		return err
	}
	// 删除角色
	if _, err := session.Table("admin_sys_role").In("id", params.Ids).Delete(); err != nil {
		return err
	}
	if err := session.Commit(); err != nil {
		return err
	}
	if err := roleService.authCache.UpdateAllRole(); err != nil {
		return err
	}
	if err := roleService.authCache.UpdateAllPerm(); err != nil {
		return err
	}
	return nil
}
