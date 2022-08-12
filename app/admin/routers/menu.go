package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useMenu() {
	router := admin.router.Group("menu").Use(middlewares.JwtAuth()).Use(middlewares.OperationRecorder())
	{
		router.GET("/permMenu", admin.Menu.PermMenu)
		router.GET("/list", auth.Perms([]string{"sys:menu:list"}), admin.Menu.List)
		router.GET("/info", auth.Perms([]string{"sys:menu:info"}), admin.Menu.Info)
		router.POST("/add", auth.Perms([]string{"sys:menu:add"}), admin.Menu.Add)
		router.POST("/update", auth.Perms([]string{"sys:menu:update"}), admin.Menu.Update)
		router.POST("/del", auth.Perms([]string{"sys:menu:del"}), admin.Menu.Del)
	}
}
