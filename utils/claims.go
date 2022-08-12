package utils

import (
	"seed-admin/common"

	"github.com/gin-gonic/gin"
)

// 从token里解析userId
func GetUserId(ctx *gin.Context) int {
	if claims, ok := ctx.Get("claims"); !ok {
		common.LOG.Error("解析userId出错,")
		return 0
	} else {
		return claims.(*CustomerClaims).UserId
	}
}
