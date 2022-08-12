package routers

import (
	"seed-admin/common/middlewares"
)

func (admin *Admin) useCommon() {
	router := admin.router.Group("common").Use(middlewares.JwtAuth())
	{
		router.POST("/uploadImages", admin.Common.UploadImages)
	}
}
