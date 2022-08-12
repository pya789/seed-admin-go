package services

import (
	"errors"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/response"
	"seed-admin/common"
	"seed-admin/utils"
)

type DeptService struct {
	authCache utils.AuthCache
}

// 获取所有部门
func (*DeptService) GetAllDept() ([]response.DeptTree, error) {
	depts := make([]response.DeptTree, 0)
	sql := `
	SELECT 
		* 
	FROM 
		admin_sys_dept
	ORDER BY 
		sort
	`
	if err := common.DB.SQL(sql).Find(&depts); err != nil {
		return nil, err
	}
	res := deptTree(depts, 0)
	return res, nil
}

// 增加部门
func (*DeptService) AddDept(params *request.DeptAdd) error {
	dept := &entity.AdminSysDept{
		Name:     *params.Name,
		ParentId: *params.ParentId,
		Sort:     params.Sort,
	}
	if _, err := common.DB.Insert(dept); err != nil {
		return err
	}
	return nil
}

// 更新部门
func (*DeptService) UpdateDept(params *request.DeptUpdate) error {
	if *params.Id == 1 {
		if *params.ParentId != 0 {
			return errors.New("顶级部门的上级不可更改")
		}
	}
	if *params.Id == *params.ParentId {
		return errors.New("不能自己成为自己的上级")
	}
	dept := &entity.AdminSysDept{
		Name:     *params.Name,
		ParentId: *params.ParentId,
		Sort:     params.Sort,
	}
	if _, err := common.DB.Where("id = ?", params.Id).AllCols().Update(dept); err != nil {
		return err
	}
	return nil
}

// 删除部门
func (*DeptService) DelDept(params *request.DeptDel) error {
	if *params.Pid == 1 {
		return errors.New("顶级部门不可删除")
	}
	session := common.DB.NewSession()
	defer session.Close()
	// 事务开启
	if err := session.Begin(); err != nil {
		return err
	}
	pid := 0
	if ok, err := session.Table("admin_sys_dept").Where("id = ?", params.Pid).Cols("parent_id").Get(&pid); !ok {
		if err != nil {
			return err
		}
		return errors.New("获取父级的上级ID失败")
	}
	// 删除部门及子部门
	if _, err := session.Table("admin_sys_dept").In("id", params.Ids).Delete(); err != nil {
		return err
	}
	// 处理用户
	if params.UserDel {
		if _, err := session.Table("admin_sys_user").In("dept_id", params.Ids).Delete(); err != nil {
			return err
		}
	} else {
		user := new(entity.AdminSysUser)
		user.DeptId = pid
		if _, err := session.In("dept_id", params.Ids).Cols("dept_id").Update(user); err != nil {
			return err
		}
	}
	return session.Commit()
}

// 获取部门信息
func (*DeptService) GetInfo(id *int) (*entity.AdminSysDept, error) {
	dept := new(entity.AdminSysDept)
	if ok, err := common.DB.Where("id = ?", id).Get(dept); !ok {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("获取部门信息失败")
	}
	return dept, nil
}

// 根据用户ID获取部门权限
func (deptService *DeptService) GetIds(userId int) ([]int, error) {
	deptIds := make([]int, 0)
	roleIds, err := deptService.authCache.GetRoleIds(userId)
	if err != nil {
		return nil, err
	}
	if err := common.DB.Table("admin_sys_role_dept").In("role_id", roleIds).Cols("dept_id").Find(&deptIds); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return deptIds, nil
}

// 递归获取部门树
func deptTree(depts []response.DeptTree, pid int) []response.DeptTree {
	var nodes = make([]response.DeptTree, 0, len(depts))
	for _, item := range depts {
		if item.ParentId == pid {
			item.Children = append(item.Children, deptTree(depts, item.Id)...)
			nodes = append(nodes, item)
		}
	}
	return nodes
}
