package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"

	"github.com/gin-gonic/gin"
)

type Log struct{}

var logService services.LogService

func (*Log) List(ctx *gin.Context) {
	var params request.LogList
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, count, err := logService.GetAllLog(&params)
	if err != nil {
		common.LOG.Error(err.Error())
	}
	common.OkData(ctx, map[string]any{
		"list":  res,
		"count": count,
	})
}
func (*Log) Del(ctx *gin.Context) {
	var params request.LogDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := logService.DelLog(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "删除日志成功")
}
