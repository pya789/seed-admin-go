package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useDept() {
	router := admin.router.Group("dept").Use(middlewares.JwtAuth()).Use(middlewares.OperationRecorder())
	{
		router.GET("/list", auth.Perms([]string{"sys:dept:list"}), admin.Dept.List)
		router.GET("/info", auth.Perms([]string{"sys:dept:info"}), admin.Dept.Info)
		router.POST("/add", auth.Perms([]string{"sys:dept:add"}), admin.Dept.Add)
		router.POST("/update", auth.Perms([]string{"sys:dept:update"}), admin.Dept.Update)
		router.POST("/del", auth.Perms([]string{"sys:dept:del"}), admin.Dept.Del)
	}
}
