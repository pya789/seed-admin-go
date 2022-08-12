package common

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/gookit/config/v2"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

var (
	DB        *xorm.Engine
	CONFIG    *config.Config
	LOG       *zap.Logger
	Redis     *redis.Client
	StartTime time.Time
)
