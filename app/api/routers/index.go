package routers

import (
	"seed-admin/app/api/controllers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.RouterGroup
	User   *controllers.User
}

func New(router *gin.RouterGroup) {
	controllers := new(Api)
	controllers.router = router.Group("api")
	controllers.useUser()
}
