package auth

import (
	"blog-service/pkg/redis"
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func SetToken(tokenString string, userID uint, expiredAt time.Time) error {
	ctx := context.Background()
	prefix := viper.GetString("redis.prefix")
	sub_prefix := viper.GetString("redis.sub_prefix1")
	key := fmt.Sprintf("%s:%s:%d", prefix, sub_prefix, userID)
	err := redis.RDB.Set(ctx, key, tokenString, time.Until(expiredAt)).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetToken(userID uint) (string, error) {
	ctx := context.Background()
	prefix := viper.GetString("redis.prefix")
	sub_prefix := viper.GetString("redis.sub_prefix1")
	key := fmt.Sprintf("%s:%s:%d", prefix, sub_prefix, userID)
	tokenString, err := redis.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func BanToken(tokenString string, expiredAt time.Time) error {
	ctx := context.Background()
	prefix := viper.GetString("redis.prefix")
	sub_prefix := viper.GetString("redis.sub_prefix2")
	key := fmt.Sprintf("%s:%s:%s", prefix, sub_prefix, tokenString)
	err := redis.RDB.Set(ctx, key, "banned", time.Until(expiredAt)).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsBanned(tokenString string) (bool, error) {
	ctx := context.Background()
	prefix := viper.GetString("redis.prefix")
	sub_prefix := viper.GetString("redis.sub_prefix2")
	key := fmt.Sprintf("%s:%s:%s", prefix, sub_prefix, tokenString)

	exists, err := redis.RDB.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists != 0, nil
}

func HasToken(userID uint) (bool, error) {
	ctx := context.Background()
	prefix := viper.GetString("redis.prefix")
	sub_prefix := viper.GetString("redis.sub_prefix1")
	key := fmt.Sprintf("%s:%s:%d", prefix, sub_prefix, userID)
	exists, err := redis.RDB.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
