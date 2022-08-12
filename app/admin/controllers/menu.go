package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Menu struct{}

var menuService services.MenuService

// 菜单和权限
func (*Menu) PermMenu(ctx *gin.Context) {
	menus, err := menuService.GetMenu(utils.GetUserId(ctx))
	if err != nil {
		common.LOG.Error(err.Error())
	}
	perms, err := menuService.GetPerms(utils.GetUserId(ctx))
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, map[string]any{
		"menus": menus,
		"perms": perms,
	})
}

// 菜单列表
func (*Menu) List(ctx *gin.Context) {
	name := ctx.Query("name")
	status := ctx.Query("status")
	res, err := menuService.GetAllMenu(name, status)
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, res)
}

// 获取菜单信息
func (*Menu) Info(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := menuService.GetInfo(id)
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, res)
}

// 增加菜单
func (*Menu) Add(ctx *gin.Context) {
	var params request.Menu
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := menuService.AddMenu(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "增加菜单成功")
}

// 编辑菜单
func (*Menu) Update(ctx *gin.Context) {
	var params request.Menu
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	err := menuService.UpdateMenu(&params)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "更新菜单成功")
}

// 删除菜单
func (*Menu) Del(ctx *gin.Context) {
	var params request.MenuDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := menuService.DelMenu(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "删除菜单成功")
}
