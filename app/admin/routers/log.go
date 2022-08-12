package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useLog() {
	router := admin.router.Group("log").Use(middlewares.JwtAuth())
	{
		router.GET("/list", auth.Perms([]string{"sys:log:list"}), admin.Log.List)
		router.POST("/del", auth.Perms([]string{"sys:log:del"}), admin.Log.Del)
	}
}
