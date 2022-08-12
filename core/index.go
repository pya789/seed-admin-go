package core

import (
	"fmt"
	"net/http"
	routers "seed-admin/app"
	common "seed-admin/common"
	config "seed-admin/core/configuration"
	middlewares "seed-admin/core/middlewares"
	redis "seed-admin/core/redis"
	xorm "seed-admin/core/xorm"
	log "seed-admin/core/zap"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化
func Run() {
	common.CONFIG = config.AddConfig()
	common.LOG = log.AddZap()
	common.DB = xorm.AddXorm()
	common.Redis = redis.AddGoRedis()
	// 应用启动时间
	common.StartTime = time.Now()
	// 运行环境
	gin.SetMode(common.CONFIG.String("app.mode"))
	app := gin.New()
	// 加载全局中间件
	middlewares.Load(app)
	// 加载路由
	routers.Load(app)
	// 静态目录
	app.StaticFS(common.CONFIG.String("app.staticPath"), http.Dir("."+common.CONFIG.String("app.staticPath")))
	// 广告
	fmt.Println(`
		欢迎使用 Seed-Admin
		当前版本:V1.0.0
	    QQ交流1群：8455822
	`)
	// 冲
	app.Run(":" + common.CONFIG.String("app.port"))
}
