package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useServer() {
	router := admin.router.Group("server").Use(middlewares.JwtAuth())
	{
		router.GET("/info", auth.Perms([]string{"sys:server:info"}), admin.Server.Info)
	}
}
