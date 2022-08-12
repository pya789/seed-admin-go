package auth

import (
	"errors"
	"seed-admin/common"
	"seed-admin/utils"

	"github.com/gin-gonic/gin"
)

const (
	// OR逻辑常量
	OR = "or"
	// AND逻辑常量
	AND = "and"
)

// 验证角色
// roleLabels：要验证的角色label
// logical：auth.OR || auth.AND 默认OR
func Roles(roleLabels []string, logical ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := utils.GetUserId(ctx)
		AuthCache := new(utils.AuthCache)
		labelsCache, err := AuthCache.GetRoleLabels(userId)
		if err != nil {
			common.FailMsg(ctx, err.Error())
			ctx.Abort()
			return
		}
		if len(logical) > 0 {
			switch logical[0] {
			case "or":
				if err := orVerify(labelsCache, roleLabels); err != nil {
					common.FailMsg(ctx, "未满足接口所需角色")
					ctx.Abort()
					return
				}
			case "and":
				if err := andVerify(labelsCache, roleLabels); err != nil {
					common.FailMsg(ctx, "未满足接口所需角色")
					ctx.Abort()
					return
				}
			default:
				common.FailMsg(ctx, "角色验证参数错误")
				ctx.Abort()
				return
			}
		} else {
			// 默认的验证条件 需要默认OR可以把函数改成orVerify
			if err := andVerify(labelsCache, roleLabels); err != nil {
				common.FailMsg(ctx, "未满足接口所需角色")
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}

// 验证权限
// menuPerms：要验证的菜单权限
// logical：auth.OR || auth.AND (默认AND)
func Perms(menuPerms []string, logical ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := utils.GetUserId(ctx)
		AuthCache := new(utils.AuthCache)
		permsCache, err := AuthCache.GetMenuPerms(userId)
		if err != nil {
			common.FailMsg(ctx, err.Error())
			ctx.Abort()
			return
		}
		if len(logical) > 0 {
			switch logical[0] {
			case "or":
				if err := orVerify(permsCache, menuPerms); err != nil {
					common.FailMsg(ctx, "权限验证失败")
					ctx.Abort()
					return
				}
			case "and":
				if err := andVerify(permsCache, menuPerms); err != nil {
					common.FailMsg(ctx, "权限验证失败")
					ctx.Abort()
					return
				}
			default:
				common.FailMsg(ctx, "权限验证参数错误")
				ctx.Abort()
				return
			}

		} else {
			// 默认的验证条件 需要默认OR可以把函数改成orVerify
			if err := andVerify(permsCache, menuPerms); err != nil {
				common.FailMsg(ctx, "权限验证失败")
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}

// OR验证
func orVerify(cacheData []string, verifyData []string) error {
	is := false
	for _, perm := range verifyData {
		if utils.SliceIncludes(cacheData, perm) {
			is = true
		}
	}
	if !is {
		return errors.New("验证失败")
	}
	return nil
}

// AND验证
func andVerify(cacheData []string, verifyData []string) error {
	count := 0
	for _, perm := range cacheData {
		if utils.SliceIncludes(verifyData, perm) {
			count++
		}
	}
	if count != len(verifyData) {
		return errors.New("验证失败")
	}
	return nil
}
