package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 没有密码则留空
		DB:       1,                // 使用数据库
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("无法连接到 Redis:", err)
		return
	}
	fmt.Println("连接成功:", pong)
	defer rdb.Close()

	articles := getAllArticles()

	var itoid = make(map[int]uint)
	var idtoi = make(map[uint]int)
	// redis中维护一个访问次数的有序集合
	for i := 0; i < len(articles); i++ {
		rdb.ZAdd(ctx, "visited", redis.Z{
			Score:  float64(articles[i].Visited),
			Member: articles[i].ID,
		})
		itoid[i] = articles[i].ID
		idtoi[articles[i].ID] = i
	}

	// 缓存前n篇文章
	articleIDsToCache, _ := rdb.ZRevRange(ctx, "visited", 0, 10).Result()
	// fmt.Println(articlesToCache)
	for _, idstr := range articleIDsToCache {
		id, _ := strconv.Atoi(idstr)
		rdb.HSet(ctx, "articles:"+string(id), articles[idtoi[uint(id)]])
	}

	// 定时任务自动同步
	go autoSyncVisited()
	// 退出前手动同步一次
	defer syncVisited()

}

// 定时同步数据到数据库
func autoSyncVisited() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			syncVisited()
		}
	}()
	select {}
}

// 进行一次同步
func syncVisited() {
	var cursor uint64
	for {
		// 获取成员
		var idstrs []string
		idstrs, cursor, _ = rdb.ZScan(ctx, "visited", 0, "*", 10).Result()

		//遍历成员
		for _, idstr := range idstrs {
			//获取id
			id, _ := strconv.Atoi(idstr)
			//获取访问次数
			visitedfloat, _ := rdb.ZScore(ctx, "visited", "idstr").Result()
			// 更新访问次数
			updateVisited(uint(id), int(visitedfloat))
		}

		if cursor == 0 {
			break
		}
	}
}
