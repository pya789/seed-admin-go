package controllers

import (
	"seed-admin/app/admin/services"
	"seed-admin/common"

	"github.com/gin-gonic/gin"
)

type Common struct{}

var commonService services.CommonService

// 上传图片
func (*Common) UploadImages(ctx *gin.Context) {
	res, err := commonService.Uploads(ctx)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsgData(ctx, "上传成功", res)
}
