package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"

	"github.com/gin-gonic/gin"
)

type Dept struct{}

var deptService services.DeptService

func (*Dept) List(ctx *gin.Context) {
	res, err := deptService.GetAllDept()
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, res)
}

// 增加部门
func (*Dept) Add(ctx *gin.Context) {
	var params request.DeptAdd
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := deptService.AddDept(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "增加部门成功")
}

// 更新部门
func (*Dept) Update(ctx *gin.Context) {
	var params request.DeptUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := deptService.UpdateDept(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "更新部门成功")
}

// 删除部门
func (*Dept) Del(ctx *gin.Context) {
	var params request.DeptDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := deptService.DelDept(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "删除部门成功")
}

// 部门信息
func (*Dept) Info(ctx *gin.Context) {
	var params request.DeptInfo
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, err := deptService.GetInfo(params.Id)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, res)
}
