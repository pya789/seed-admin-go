package routers

import (
	"seed-admin/common/auth"
	"seed-admin/common/middlewares"
)

func (admin *Admin) useDict() {
	router := admin.router.Group("dict").Use(middlewares.JwtAuth()).Use(middlewares.OperationRecorder())
	{
		router.GET("/list", auth.Perms([]string{"sys:dict:list"}), admin.Dict.List)
		router.GET("/info", auth.Perms([]string{"sys:dict:info"}), admin.Dict.Info)
		router.POST("/add", auth.Perms([]string{"sys:dict:add"}), admin.Dict.Add)
		router.POST("/update", auth.Perms([]string{"sys:dict:update"}), admin.Dict.Update)
		router.POST("/del", auth.Perms([]string{"sys:dict:del"}), admin.Dict.Del)
		router.GET("/dataList", auth.Perms([]string{"sys:dict:details:list"}), admin.Dict.DataList)
		router.GET("/dataInfo", auth.Perms([]string{"sys:dict:details:info"}), admin.Dict.DataInfo)
		router.POST("/dataAdd", auth.Perms([]string{"sys:dict:details:add"}), admin.Dict.DataAdd)
		router.POST("/dataUpdate", auth.Perms([]string{"sys:dict:details:update"}), admin.Dict.DataUpdate)
		router.POST("/dataDel", auth.Perms([]string{"sys:dict:details:del"}), admin.Dict.DataDel)
		router.GET("/typeData", admin.Dict.TypeData)
	}
}
