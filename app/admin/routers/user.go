package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useUser() {
	router := admin.router.Group("user")
	{
		router.POST("/login", middlewares.OperationRecorder(), admin.User.Login)
		router.GET("/captcha", admin.User.Captcha)
		jwtAuth := router.Group("").Use(middlewares.JwtAuth(), middlewares.OperationRecorder())
		{
			jwtAuth.GET("/person", admin.User.Person)
			jwtAuth.POST("/list", auth.Perms([]string{"sys:user:list"}), admin.User.List)
			jwtAuth.POST("/add", auth.Perms([]string{"sys:user:add"}), admin.User.Add)
			jwtAuth.POST("/del", auth.Perms([]string{"sys:user:del"}), admin.User.Del)
			jwtAuth.GET("/info", auth.Perms([]string{"sys:user:info"}), admin.User.Info)
			jwtAuth.POST("/update", auth.Perms([]string{"sys:user:update"}), admin.User.Update)
			jwtAuth.POST("/move", auth.Perms([]string{"sys:user:move"}), admin.User.Move)
			jwtAuth.POST("/updateUserRole", auth.Perms([]string{"sys:user:updateUserRole"}), admin.User.UpdateUserRole)
			jwtAuth.POST("/updateAvatar", admin.User.UpdateAvatar)
			jwtAuth.POST("/updateBaseInfo", admin.User.UpdateBaseInfo)
			jwtAuth.POST("/updatePassword", admin.User.UpdatePassword)
		}
	}
}
