package controllers

import (
	"seed-admin/common"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Demo(ctx *gin.Context) {
	common.OkMsg(ctx, "我只是个demo啊")
}
