package redis

import (
	"seed-admin/common"

	"github.com/go-redis/redis"
)

func AddGoRedis() *redis.Client {
	address := common.CONFIG.String("redis.server") + ":" + common.CONFIG.String("redis.port")
	password := common.CONFIG.String("redis.password")
	db := common.CONFIG.Int("redis.database")
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	if _, err := client.Ping().Result(); err != nil {
		common.LOG.Error(err.Error())
	}
	return client
}
