package redis

import (
	"blog-service/pkg/logger"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	})

	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Error("Error connecting to Redis", zap.Error(err))
		panic(err)
	}
	logger.Logger.Info("Successfully connected to Redis")
}
