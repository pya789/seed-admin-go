package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"seed-admin/app/admin/entity"
	"seed-admin/app/admin/services"
	"seed-admin/common"
	"seed-admin/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var logService services.LogService

// 操作记录器
func OperationRecorder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := make([]byte, 0)
		// 只考虑常规的浏览器取参方式 get携带body post携带url等不再进行取参处理 有需要自己可以添加
		// 本项目不遵守Restful规范 delete,put等请求方法如有需要请自行添加case分支 这里只处理GETQuery和POSTBody的请求方式
		switch ctx.Request.Method {
		case http.MethodGet:
			body = []byte(ctx.Request.URL.Query().Encode())
		case http.MethodPost:
			var err error
			body, err = io.ReadAll(ctx.Request.Body)
			if err != nil {
				common.LOG.Error(err.Error())
			} else {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}
		userId := 0
		if claims, ok := ctx.Get("claims"); ok {
			userId = claims.(*utils.CustomerClaims).UserId
		}
		// 5分钟内重复操作直接next不再重复记录
		if userId != 0 {
			ok, err := common.Redis.SIsMember("record"+strconv.Itoa(userId), ctx.Request.URL.Path).Result()
			if err != nil {
				common.LOG.Error(err.Error())
			}
			if ok {
				ctx.Next()
				return
			}
		}
		log := &entity.AdminSysLog{
			UserId: userId,
			Method: ctx.Request.Method,
			Action: ctx.Request.URL.Path,
			Ip:     ctx.ClientIP(),
			Params: string(body),
		}
		writer := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = writer
		ctx.Next()
		log.StatusCode = ctx.Writer.Status()
		log.Results = writer.body.String()
		if err := logService.AddLog(log); err != nil {
			common.LOG.Error(err.Error())
		}
		// 把操作用redis记录下来
		if userId != 0 {
			common.Redis.SAdd("record"+strconv.Itoa(userId), ctx.Request.URL.Path).Result()
			common.Redis.Expire("record"+strconv.Itoa(userId), time.Second*300).Result()
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
