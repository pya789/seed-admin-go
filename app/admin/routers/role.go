package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useRole() {
	router := admin.router.Group("role").Use(middlewares.JwtAuth()).Use(middlewares.OperationRecorder())
	{
		router.GET("/list", auth.Perms([]string{"sys:role:list"}), admin.Role.List)
		router.GET("/info", auth.Perms([]string{"sys:role:info"}), admin.Role.Info)
		router.POST("/add", auth.Perms([]string{"sys:role:add"}), admin.Role.Add)
		router.POST("/update", auth.Perms([]string{"sys:role:update"}), admin.Role.Update)
		router.POST("/del", auth.Perms([]string{"sys:role:del"}), admin.Role.Del)
	}
}
