package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

func setSessionID(userID uint, sessionID string) {
	key := fmt.Sprintf("sid:%d", userID)
	rdb.Set(context.Background(), key, sessionID, 1*time.Hour)
	key = fmt.Sprintf("uid:%s", sessionID)
	rdb.Set(context.Background(), key, userID, 1*time.Hour)
}

func deleteSessionID(sessionID string) {
	userID := getUserID(sessionID)
	key := fmt.Sprintf("uid:%s", sessionID)
	rdb.Del(context.Background(), key)
	key = fmt.Sprintf("sid:%d", userID)
	rdb.Del(context.Background(), key)
}

// func getSessionID(userID uint) (sessionID string) {
// 	key := fmt.Sprintf("sid:%d", userID)
// 	sessionID = rdb.Get(context.Background(), key).String()
// 	return
// }

func getUserID(sessionID string) (userID uint) {
	key := fmt.Sprintf("uid:%s", sessionID)
	userIDstr := rdb.Get(context.Background(), key).Val()
	userIDint, _ := strconv.Atoi(userIDstr)
	return uint(userIDint)
}
