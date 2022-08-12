package routers

import (
	"seed-admin/app/admin/controllers"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	router *gin.RouterGroup
	User   *controllers.User
	Menu   *controllers.Menu
	Role   *controllers.Role
	Dept   *controllers.Dept
	Dict   *controllers.Dict
	Log    *controllers.Log
	Server *controllers.Server
	Common *controllers.Common
}

func New(router *gin.RouterGroup) {
	controllers := new(Admin)
	controllers.router = router.Group("admin")
	controllers.useUser()
	controllers.useMenu()
	controllers.useRole()
	controllers.useDept()
	controllers.useDict()
	controllers.useLog()
	controllers.useServer()
	controllers.useCommon()
}
