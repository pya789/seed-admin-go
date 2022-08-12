package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS            = 1000 // 通用成功
	FAIL               = 1001 // 通用失败 前端会自动抛出异常
	REFRESH_CAPTCHA    = 1002 // 需要前端手动判断code == 1002处理的失败
	AUTHORIZATION_FAIL = 1004 // 鉴权失败
)

// 自定义通用消息
func Message(c *gin.Context, status int, message string, data ...any) {
	var obj gin.H
	if len(data) == 0 {
		obj = gin.H{
			"code":    status,
			"message": message,
		}
	} else {
		obj = gin.H{
			"code":    status,
			"message": message,
			"data":    data[0],
		}
	}
	c.JSON(http.StatusOK, obj)
}

// 默认的成功响应
func Ok(c *gin.Context) {
	obj := gin.H{
		"code":    SUCCESS,
		"message": "操作成功",
	}
	c.JSON(http.StatusOK, obj)
}

// 携带消息的成功响应
func OkMsg(c *gin.Context, message string) {
	obj := gin.H{
		"code":    SUCCESS,
		"message": message,
	}
	c.JSON(http.StatusOK, obj)
}

// 携带数据的成功响应
func OkData(c *gin.Context, data any) {
	obj := gin.H{
		"code":    SUCCESS,
		"message": "操作成功",
		"data":    data,
	}
	c.JSON(http.StatusOK, obj)
}

// 携带消息和数据的成功响应
func OkMsgData(c *gin.Context, message string, data any) {
	obj := gin.H{
		"code":    SUCCESS,
		"message": message,
		"data":    data,
	}
	c.JSON(http.StatusOK, obj)
}

// 默认的失败响应
func Fail(c *gin.Context) {
	obj := gin.H{
		"code":    FAIL,
		"message": "操作失败",
	}
	c.JSON(http.StatusOK, obj)
}

// 携带消息的失败响应
func FailMsg(c *gin.Context, message string) {
	obj := gin.H{
		"code":    FAIL,
		"message": message,
	}
	c.JSON(http.StatusOK, obj)
}

// 携带数据的失败响应
func FailData(c *gin.Context, data any) {
	obj := gin.H{
		"code":    FAIL,
		"message": "操作失败",
		"data":    data,
	}
	c.JSON(http.StatusOK, obj)
}

// 携带消息和数据的失败响应
func FailMsgData(c *gin.Context, message string, data any) {
	obj := gin.H{
		"code":    FAIL,
		"message": message,
		"data":    data,
	}
	c.JSON(http.StatusOK, obj)
}
