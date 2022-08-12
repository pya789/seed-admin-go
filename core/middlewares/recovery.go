package middlewares

import "github.com/gin-gonic/gin"

func useRecovery(r *gin.Engine) {
	// 加载官方的恢复 如有必要自己写
	r.Use(gin.Recovery())
}
