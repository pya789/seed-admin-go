package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"

	"github.com/gin-gonic/gin"
)

type Dict struct{}

var dictService services.DictService

// 字典列表
func (*Dict) List(ctx *gin.Context) {
	var params request.DictList
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, count, err := dictService.GetAllDictType(&params)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, map[string]any{
		"list":  res,
		"count": count,
	})
}

// 增加字典
func (*Dict) Add(ctx *gin.Context) {
	var params request.DictAdd
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.AddDictType(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "增加字典成功")
}
func (*Dict) Update(ctx *gin.Context) {
	var params request.DictUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.UpdateDictType(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "更新字典成功")
}

// 删除字典
func (*Dict) Del(ctx *gin.Context) {
	var params request.DictDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.DelDictType(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "删除字典成功")
}

// 获取字典信息
func (*Dict) Info(ctx *gin.Context) {
	var params request.DictInfo
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, err := dictService.GetDictTypeInfo(params.Id)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, res)
}

// 字典数据列表
func (*Dict) DataList(ctx *gin.Context) {
	var params request.DictDataList
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, count, err := dictService.GetAllDictData(&params)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, map[string]any{
		"list":  res,
		"count": count,
	})
}

// 增加字典数据
func (*Dict) DataAdd(ctx *gin.Context) {
	var params request.DictDataAdd
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.AddDictData(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "增加字典数据成功")
}

// 字典数据更新
func (*Dict) DataUpdate(ctx *gin.Context) {
	var params request.DictDataUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.UpdateDictData(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "更新字典数据成功")
}

// 删除字典数据
func (*Dict) DataDel(ctx *gin.Context) {
	var params request.DictDataDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := dictService.DelDictData(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
	}
	common.OkMsg(ctx, "删除字典数据成功")
}

// 获取字典数据信息
func (*Dict) DataInfo(ctx *gin.Context) {
	var params request.DictDataInfo
	_ = ctx.ShouldBindQuery(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, err := dictService.GetDictDataInfo(params.Id)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, res)
}

// 根据类型获取字典
func (*Dict) TypeData(ctx *gin.Context) {
	dictType := ctx.Query("type")
	res, err := dictService.GetTypeData(dictType)
	if err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, res)
}
