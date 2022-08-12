package controllers

import (
	"seed-admin/app/admin/services"
	"seed-admin/common"

	"github.com/gin-gonic/gin"
)

type Server struct{}

var serverService services.ServerService

// 服务器信息
func (*Server) Info(ctx *gin.Context) {
	res, err := serverService.GetSystemInfo()
	if err != nil {
		common.FailMsg(ctx, "获取服务器信息失败")
		return
	}
	common.OkData(ctx, res)
}
