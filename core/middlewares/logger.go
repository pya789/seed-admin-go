package middlewares

import (
	"seed-admin/common"

	"seed-admin/core/zap/lumberjack"

	"github.com/gin-gonic/gin"
)

// // zap接管gin日志(有需要自行加载)
// func UseZapLogger(r *gin.Engine) {
// 	logger := common.LOG
// 	r.Use(func(ctx *gin.Context) {
// 		start := time.Now()
// 		ctx.Next()
// 		cost := time.Since(start)
// 		logger.Info(ctx.Request.URL.Path,
// 			zap.Int("status", ctx.Writer.Status()),
// 			zap.String("method", ctx.Request.Method),
// 			zap.String("path", ctx.Request.URL.Path),
// 			zap.String("query", ctx.Request.URL.RawQuery),
// 			zap.String("ip", ctx.ClientIP()),
// 			zap.String("user-agent", ctx.Request.UserAgent()),
// 			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
// 			zap.Duration("cost", cost),
// 		)
// 	})
// }
func UseLogger(r *gin.Engine) {
	var loggerConfig = gin.LoggerConfig{
		Output: &lumberjack.Logger{
			Filename:   "./" + common.CONFIG.String("gin_log.director") + "/http.log",
			MaxSize:    common.CONFIG.Int("gin_log.maxSize"),
			MaxBackups: common.CONFIG.Int("gin_log.maxBackups"),
			MaxAge:     common.CONFIG.Int("gin_log.maxAge"),
			Compress:   common.CONFIG.Bool("gin_log.compress"),
		},
	}
	switch common.CONFIG.String("gin_log.outType") {
	case "console":
		r.Use(gin.Logger())
	case "file":
		r.Use(gin.LoggerWithConfig(loggerConfig))
	default:
		r.Use(gin.Logger()).Use(gin.LoggerWithConfig(loggerConfig))
	}
}
