package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// Message 结构体表示复合消息
type Message struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Msg      string `json:"msg"`
	Channel  string `json:"channel"`
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       1,  // 默认DB 0
	})

	go func() {
		channels := []string{"ch1", "ch2", "ch3", "ch4", "ch5"}
		// 开启五个频道
		rdb.Subscribe(context.Background(), channels...)
	}()
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

func sendMessage(message Message) {
	messmageJson, _ := json.Marshal(message)
	rdb.Publish(context.Background(), message.Channel, messmageJson)
}

func subAllChannels() <-chan *redis.Message {
	channels := getChannelsList()
	ch := rdb.PSubscribe(context.Background(), channels...).Channel()
	return ch
}
func getChannelsList() []string {
	channels, _ := rdb.PubSubChannels(context.Background(), "*").Result()
	return channels
}
