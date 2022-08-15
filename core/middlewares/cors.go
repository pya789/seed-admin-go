package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UseCors(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		// 允许来源
		c.Header("Access-Control-Allow-Origin", origin)
		// 请求方式
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 允许的标头
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		// 暴露标头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// 是否允许携带cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		// 预检时间(秒)
		c.Header("Access-Control-Max-Age", "3600")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})
}
