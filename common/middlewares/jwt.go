package middlewares

import (
	"seed-admin/common"
	"seed-admin/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" || tokenString == "Bearer " {
			common.Message(ctx, common.AUTHORIZATION_FAIL, "未携带token,认证失败")
			ctx.Abort()
			return
		}
		jwt := utils.NewJwt()
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			common.Message(ctx, common.AUTHORIZATION_FAIL, "非法token或token已经过期,请重新登录")
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
