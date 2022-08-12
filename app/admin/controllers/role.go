package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Role struct{}

var roleService services.RoleService

// 获取角色列表
func (*Role) List(ctx *gin.Context) {
	var params request.RoleList
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, count, err := roleService.GetAllRole(&params)
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, map[string]any{
		"list":  res,
		"count": count,
	})
}

// 获取角色信息
func (*Role) Info(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := roleService.GetInfo(id)
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, res)
}

// 新增角色
func (*Role) Add(ctx *gin.Context) {
	var params request.Role
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := roleService.AddRole(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.Ok(ctx)
}

// 编辑角色
func (*Role) Update(ctx *gin.Context) {
	var params request.RoleUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := roleService.UpdateRole(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "更新角色成功")
}

// 删除角色
func (*Role) Del(ctx *gin.Context) {
	var params request.RoleDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := roleService.DelRole(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "删除角色成功")
}
