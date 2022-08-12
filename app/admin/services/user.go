package services

import (
	"errors"
	"fmt"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/response"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService struct {
	authCache   utils.AuthCache
	menuService MenuService
	deptService DeptService
}

// 登录
func (*UserService) Login(params *request.Login) (*entity.AdminSysUser, []int, error) {
	// MD5加盐 盐值一旦设定并在生产模式使用后切勿更改 随意更改后会造成之前盐值不同的用户无法登录
	params.Password = utils.Md5Salt(params.Password, common.CONFIG.String("app.md5Salt"))
	// 使用帐密做查询
	user := &entity.AdminSysUser{
		Username: params.Username,
		Password: params.Password,
	}
	if ok, err := common.DB.Get(user); ok {
		if user.Status == 1 {
			return nil, nil, errors.New("该用户已被禁用")
		}
		// 查询用户拥有的角色ID
		roleIds := make([]int, 0)
		if err = common.DB.Table("admin_sys_user_role").
			Where("user_id = ?", user.Id).Cols("role_id").
			Find(&roleIds); err != nil {
			return nil, nil, err
		}
		return user, roleIds, nil
	} else {
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("帐号或密码错误")
	}
}

// 生成一个token
func (*UserService) GetToken(user *entity.AdminSysUser) (string, error) {
	j := utils.NewJwt()
	token, err := j.CreateToken(utils.CustomerClaims{
		UserId: user.Id,
		StandardClaims: &jwt.StandardClaims{
			Issuer:    common.CONFIG.String("jwt.Issuer"),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + common.CONFIG.Int64("jwt.ExpireSeconds"),
		},
	})
	return token, err
}

// 获取个人用户信息
func (userService *UserService) GetPerson(userId int) (*response.UserProfile, error) {
	user := new(response.UserProfile)
	if ok, err := common.DB.
		Table("admin_sys_user").
		Alias("u").
		Join("INNER", []string{"admin_sys_dept", "d"}, "u.dept_id = d.id").
		Where("u.status = ? AND u.id = ?", 0, userId).
		Get(user); ok {
		// 查询用户拥有的角色ID
		roleMap := make([]map[string]any, 0)
		if err = common.DB.Table("admin_sys_user_role").Alias("ur").
			Join("INNER", []string{"admin_sys_role", "r"}, "ur.role_id = r.id").
			Where("user_id = ?", user.Id).Cols("r.id", "r.label", "r.name").
			Find(&roleMap); err != nil {
			return nil, err
		}
		if len(roleMap) <= 0 {
			return nil, errors.New("该用户没有权限")
		}
		roleIds := make([]int, 0, len(roleMap))
		roleLabels := make([]string, 0, len(roleMap))
		roleNames := make([]string, 0, len(roleMap))
		for _, item := range roleMap {
			roleIds = append(roleIds, int(item["id"].(int32)))
			roleLabels = append(roleLabels, item["label"].(string))
			roleNames = append(roleNames, item["name"].(string))
		}
		userService.authCache.SetRoleIds(userId, roleIds)
		userService.authCache.SetRoleLabels(userId, roleLabels)
		user.Roles = roleLabels
		user.RoleNames = roleNames
		return user, nil
	} else {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("获取用户信息失败/用户被禁止使用")
	}
}

// 获取用户信息
func (*UserService) GetInfo(userId int) (*response.UserInfo, error) {
	user := &entity.AdminSysUser{
		Id: userId,
	}
	if ok, err := common.DB.Omit("password").Get(user); ok {
		// 查询用户拥有的角色ID
		roleIds := make([]int, 0)
		if err = common.DB.Table("admin_sys_user_role").
			Where("user_id = ?", user.Id).Cols("role_id").
			Find(&roleIds); err != nil {
			return nil, err
		}
		res := &response.UserInfo{
			Id:       user.Id,
			DeptId:   user.DeptId,
			Username: user.Username,
			NickName: user.NickName,
			Phone:    user.Phone,
			Email:    user.Email,
			Avatar:   user.Avatar,
			Status:   user.Status,
			RoleIds:  roleIds,
		}
		return res, nil
	} else {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("获取用户信息失败")
	}
}

// 更新头像
func (*UserService) UpdateAvatar(params *request.UserAvatarUpdate, userId int) error {
	user := &entity.AdminSysUser{
		Avatar: params.Url,
	}
	if _, err := common.DB.Where("id = ?", userId).Cols("avatar").Update(user); err != nil {
		return err
	}
	return nil
}

// 获取全部用户
func (userService *UserService) GetAllUser(params *request.UserList, userId int) ([]response.UserList, int64, error) {
	user := new(entity.AdminSysUser)
	users := make([]response.UserList, 0)
	// 获取当前操作用户的部门权限ID组
	userDeptIds, err := userService.deptService.GetIds(userId)
	if err != nil {
		return nil, 0, err
	}
	// 处理用户部门权限条件
	userDeptInStr := ""
	userDeptCountInStr := ""

	if len(userDeptIds) != 0 {
		userDeptIdsIn := utils.SliceToInStr(userDeptIds)
		userDeptInStr = fmt.Sprintf("AND user.dept_id in(%v)", userDeptIdsIn)
		userDeptCountInStr = fmt.Sprintf("dept_id in(%v)", userDeptIdsIn)
	}
	deptInStr := ""
	deptCountInStr := ""
	if len(params.DeptId) != 0 {
		deptIdsIn := utils.SliceToInStr(params.DeptId)
		deptInStr = fmt.Sprintf("AND user.dept_id in(%v)", deptIdsIn)
		deptCountInStr = fmt.Sprintf("dept_id in(%v)", deptIdsIn)
	}
	status := ""
	if params.Status != nil {
		status = strconv.Itoa(*params.Status)
	}
	count, err := common.DB.
		Where("username LIKE ?", "%"+params.Username+"%").
		Where("nick_name LIKE ?", "%"+params.NickName+"%").
		Where("phone LIKE ?", "%"+params.Phone+"%").
		Where("status LIKE ?", "%"+status+"%").
		Where(deptCountInStr).
		Where(userDeptCountInStr).
		Count(user)
	if err != nil {
		return nil, 0, err
	}
	sql := fmt.Sprintf(`
	SELECT 
		user.*,
		GROUP_CONCAT(role.id) AS role_ids,
		dept.name AS dept_name
	FROM
		admin_sys_user user 
		LEFT JOIN admin_sys_user_role user_role ON user.id = user_role.user_id 
		LEFT JOIN admin_sys_role role ON user_role.role_id = role.id
		LEFT JOIN admin_sys_dept dept ON user.dept_id = dept.id
	WHERE
		user.username LIKE '%s' AND
		user.nick_name LIKE '%s' AND
		user.phone LIKE '%s' AND
		user.status LIKE '%s' 
		%v %v
	GROUP BY user.id
	LIMIT %v,%v
	`, "%"+params.Username+"%", "%"+params.NickName+"%", "%"+params.Phone+"%", "%"+status+"%", userDeptInStr, deptInStr, (*params.PageNum-1)*(*params.PageSize), *params.PageSize)
	if err := common.DB.SQL(sql).Find(&users); err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

// 更新用户角色
func (userService *UserService) UpdateUserRole(params *request.UserRoleUpdate) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// 删除用户原有角色
	if _, err := session.Table("admin_sys_user_role").Where("user_id = ?", params.Id).Delete(); err != nil {
		return err
	}
	if len(params.RoleIds) != 0 {
		// 给用户增加新的角色
		userRole := make([]entity.AdminSysUserRole, 0, len(params.RoleIds))
		for _, roleId := range params.RoleIds {
			userRole = append(userRole, entity.AdminSysUserRole{
				UserId: *params.Id,
				RoleId: roleId,
			})
		}
		if _, err := session.Insert(userRole); err != nil {
			return err
		}
	}
	if err := session.Commit(); err != nil {
		return err
	}
	// 处理缓存权限
	role := make([]entity.AdminSysRole, 0, len(params.RoleIds))
	if err := common.DB.In("id", params.RoleIds).Find(&role); err != nil {
		return err
	}
	userService.authCache.SetRoleIds(*params.Id, params.RoleIds)
	roleLabels := make([]string, 0, len(role))
	for _, item := range role {
		roleLabels = append(roleLabels, item.Label)
	}
	userService.authCache.SetRoleLabels(*params.Id, roleLabels)
	perms, err := userService.menuService.GetPerms(*params.Id)
	if err != nil {
		return err
	}
	userService.authCache.SetMenuPerms(*params.Id, perms)
	return nil
}

// 新增用户
func (*UserService) AddUser(params *request.User) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	user := &entity.AdminSysUser{
		Username: params.Username,
		NickName: *params.NickName,
		Password: utils.Md5Salt(params.Password, common.CONFIG.String("app.md5Salt")),
		DeptId:   params.DeptId,
		Phone:    params.Phone,
		Email:    params.Email,
		Avatar:   params.Avatar,
		Status:   params.Status,
	}
	if _, err := session.Insert(user); err != nil {
		return err
	}
	if len(params.RoleIds) == 0 {
		return session.Commit()
	}
	// 增加新的角色权限
	userRole := make([]entity.AdminSysUserRole, 0, len(params.RoleIds))
	for _, roleId := range params.RoleIds {
		userRole = append(userRole, entity.AdminSysUserRole{
			UserId: user.Id,
			RoleId: roleId,
		})
	}
	if _, err := session.Insert(userRole); err != nil {
		return err
	}
	return session.Commit()
}

// 更新用户基础信息
func (*UserService) UpdateBaseInfo(params *request.UserBaseInfoUpdate, userId int) error {
	// 更新用户表
	user := &entity.AdminSysUser{
		NickName: *params.NickName,
		Phone:    params.Phone,
		Email:    params.Email,
	}
	if _, err := common.DB.Where("id = ?", userId).Cols("nick_name", "phone", "email").Update(user); err != nil {
		return err
	}
	return nil
}

// 更新用户密码
func (*UserService) UpdatePassword(params *request.UserPasswordUpdate, userId int) error {
	user := new(entity.AdminSysUser)
	if ok, err := common.DB.Where("id = ?", userId).Cols("password").Get(user); !ok {
		if err != nil {
			return err
		}
		return errors.New("获取旧密码失败")
	}
	if user.Password != utils.Md5Salt(params.OldPassword, common.CONFIG.String("app.md5Salt")) {
		return errors.New("旧密码错误,请检查旧密码是否输入正确")
	}
	user.Password = utils.Md5Salt(params.NewPassword, common.CONFIG.String("app.md5Salt"))
	// 更新用户表
	if _, err := common.DB.Where("id = ?", userId).Cols("password").Update(user); err != nil {
		return err
	}
	return nil
}

// 更新用户
func (userService *UserService) UpdateUser(params *request.UserUpdate) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// 更新用户表
	user := &entity.AdminSysUser{
		Username: params.Username,
		NickName: *params.NickName,
		Password: utils.Md5Salt(params.Password, common.CONFIG.String("app.md5Salt")),
		DeptId:   params.DeptId,
		Phone:    params.Phone,
		Email:    params.Email,
		Avatar:   params.Avatar,
		Status:   params.Status,
	}
	if _, err := session.Where("id = ?", params.Id).AllCols().Update(user); err != nil {
		return err
	}
	userRoleUpdate := &request.UserRoleUpdate{
		Id:      params.Id,
		RoleIds: params.RoleIds,
	}
	if err := userService.UpdateUserRole(userRoleUpdate); err != nil {
		return err
	}
	if params.Status != 0 {
		userService.authCache.Del([]int{*params.Id})
	}
	return session.Commit()
}

// 删除用户
func (userService *UserService) DelUser(params *request.UserDel) error {
	session := common.DB.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// 删除用户所有角色
	if _, err := session.Table("admin_sys_user_role").In("user_id", params.Ids).Delete(); err != nil {
		return err
	}
	// 删除用户
	if _, err := session.Table("admin_sys_user").In("id", params.Ids).Delete(); err != nil {
		return err
	}
	if err := session.Commit(); err != nil {
		return err
	}
	if err := userService.authCache.Del(params.Ids); err != nil {
		return err
	}
	return nil
}

// 移动部门
func (*UserService) MoveDept(params *request.UserMove) error {
	user := new(entity.AdminSysUser)
	user.DeptId = *params.DeptId
	if _, err := common.DB.In("id", params.Ids).Update(user); err != nil {
		return err
	}
	return nil
}
