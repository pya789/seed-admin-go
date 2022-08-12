package controllers

import (
	"seed-admin/app/admin/request"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

var userService services.UserService

// 登录
func (*User) Login(ctx *gin.Context) {
	// 绑定body参数到结构
	var params request.Login
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	// 验证验证码合法性 注意CaptchaVerify调用后验证码已销毁 所以需要让前端再次获取新的验证码
	if err := utils.CaptchaVerify(params.CaptchaId, params.Captcha); err != nil {
		common.Message(ctx, common.REFRESH_CAPTCHA, err.Error())
		return
	}
	// 调用服务登录
	user, roleIds, err := userService.Login(&params)
	if err != nil {
		common.Message(ctx, common.REFRESH_CAPTCHA, err.Error())
		return
	}
	if len(roleIds) == 0 {
		common.Message(ctx, common.REFRESH_CAPTCHA, "您没有任何角色权限,无法登录")
		return
	}
	// 获取token
	token, err := userService.GetToken(user)
	if err != nil {
		common.Message(ctx, common.REFRESH_CAPTCHA, err.Error())
		return
	}
	common.OkMsgData(ctx, "登录成功", map[string]string{
		"token": token,
	})
}

// 获取登录人的用户信息
func (*User) Person(ctx *gin.Context) {
	userInfo, err := userService.GetPerson(utils.GetUserId(ctx))
	if err != nil {
		common.Message(ctx, common.AUTHORIZATION_FAIL, err.Error())
		return
	}
	common.OkData(ctx, userInfo)
}

// 生成验证码
func (*User) Captcha(ctx *gin.Context) {
	captchaId := ctx.Query("captchaId")
	cap := utils.NewCaptcha()
	var image []byte
	var err error
	// 是否进入重载
	if captchaId != "" {
		_ = cap.Reload(captchaId)
		image, err = cap.ImageByte(captchaId)
		if err != nil {
			common.FailMsg(ctx, err.Error())
			return
		}
	} else {
		captchaId = cap.CreateImage()
		image, err = cap.ImageByte(captchaId)
		if err != nil {
			common.FailMsg(ctx, err.Error())
			return
		}
	}
	data := map[string]any{
		"id":    captchaId,
		"image": image,
	}
	common.OkData(ctx, data)
}

// 获取用户列表
func (*User) List(ctx *gin.Context) {
	var params request.UserList
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	res, count, err := userService.GetAllUser(&params, utils.GetUserId(ctx))
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

// 修改用户角色
func (*User) UpdateUserRole(ctx *gin.Context) {
	var params request.UserRoleUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.UpdateUserRole(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "更新用户角色成功")
}

// 修改用户头像
func (*User) UpdateAvatar(ctx *gin.Context) {
	var params request.UserAvatarUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.UpdateAvatar(&params, utils.GetUserId(ctx)); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "修改用户头像成功")
}

// 更新用户基础信息
func (*User) UpdateBaseInfo(ctx *gin.Context) {
	var params request.UserBaseInfoUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.UpdateBaseInfo(&params, utils.GetUserId(ctx)); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "修改基础信息成功")
}

// 更新用户密码
func (*User) UpdatePassword(ctx *gin.Context) {
	var params request.UserPasswordUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.UpdatePassword(&params, utils.GetUserId(ctx)); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "修改基础信息成功")
}

// 获取用户信息
func (*User) Info(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	userInfo, err := userService.GetInfo(id)
	if err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkData(ctx, userInfo)
}

// 新增用户
func (*User) Add(ctx *gin.Context) {
	var params request.User
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.AddUser(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "新增用户成功")
}

// 修改用户
func (*User) Update(ctx *gin.Context) {
	var params request.UserUpdate
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.UpdateUser(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "用户信息修改成功")
}

// 删除用户
func (*User) Del(ctx *gin.Context) {
	var params request.UserDel
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.DelUser(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "删除用户成功")
}

// 转移部门
func (*User) Move(ctx *gin.Context) {
	var params request.UserMove
	_ = ctx.ShouldBindJSON(&params)
	// 验证参数合法性
	if err := utils.ParamsVerify(&params); err != nil {
		common.FailMsg(ctx, err.Error())
		return
	}
	if err := userService.MoveDept(&params); err != nil {
		common.LOG.Error(err.Error())
		common.FailMsg(ctx, err.Error())
		return
	}
	common.OkMsg(ctx, "转移部门成功")
}
