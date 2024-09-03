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
	for _, idstr := range articleIDsToCache {
		id, _ := strconv.Atoi(idstr)
		rdb.HSet(ctx, "articles:"+idstr, articles[idtoi[uint(id)]])
		rdb.SAdd(ctx, "cached_ids", uint(id))
	}

	// 定时任务自动同步
	go autoSyncVisited()
	// 退出前手动同步一次
	defer syncVisited()

	go autoUpdateCachedArticles()
	defer updateCachedArticles()

}

// 定时同步数据到数据库
func autoSyncVisited() {
	ticker := time.NewTicker(10 * time.Second)
	// defer ticker.Stop()

	go func() {
		for range ticker.C {
			syncVisited()
		}
	}()
}

// 进行一次同步
func syncVisited() {
	var cursor uint64
	for {
		// 获取成员
		var res []string
		res, cursor, _ = rdb.ZScan(ctx, "visited", 0, "*", 10).Result()
		//遍历成员
		for i := 0; i < len(res); i += 2 {
			//获取id
			idint, _ := strconv.Atoi(res[i])
			//获取访问次数
			visitedfloat, _ := strconv.Atoi(res[i+1])
			// 更新访问次数
			updateVisited(uint(idint), int(visitedfloat))
		}

		if cursor == 0 {
			break
		}
	}
}

// 定时更新缓存的文章
func autoUpdateCachedArticles() {
	ticker := time.NewTicker(1 * time.Hour)
	// defer ticker.Stop()

	go func() {
		for range ticker.C {
			updateCachedArticles()
		}
	}()
}

// 进行一次更新
func updateCachedArticles() {

	// 获取访问最多的前n篇文章
	articleIDsToCache, _ := rdb.ZRevRange(ctx, "visited", 0, 10).Result()
	rdb.Del(ctx, "new_cache_ids")
	for _, idstr := range articleIDsToCache {
		id, _ := strconv.Atoi(idstr)
		rdb.SAdd(ctx, "new_cache_ids", uint(id))
	}

	newCacheIDs, _ := rdb.SDiff(ctx, "new_cache_ids", "cached_ids").Result()
	deleteIDs, _ := rdb.SDiff(ctx, "cached_ids", "new_cache_ids").Result()

	// 从缓存中删除
	for _, idstr := range deleteIDs {
		rdb.Del(ctx, "articles:"+idstr)
		rdb.SRem(ctx, "cached_ids", idstr)
	}
	// 添加到缓存
	for _, idstr := range newCacheIDs {
		rdb.Del(ctx, "articles:"+idstr)
		id, _ := strconv.Atoi(idstr)
		article := getArticle(uint(id))
		rdb.HSet(ctx, "articles:"+idstr, article)
		rdb.SAdd(ctx, "cached_ids", idstr)
	}
}

func visitArticleR(articleID uint) {
	rdb.ZIncrBy(ctx, "visited", 1.0, strconv.Itoa(int(articleID)))
}

func createArticleR(articleID uint) {
	rdb.ZAdd(ctx, "visited", redis.Z{
		Score:  0,
		Member: articleID,
	})
}

func deleteArticleR(articleID uint) {
	rdb.ZRem(ctx, "visited", strconv.Itoa(int(articleID)))
	if isCached(articleID) {
		rdb.SRem(ctx, "cached_ids", articleID)
		rdb.Del(ctx, "articles:"+strconv.Itoa(int(articleID)))
		updateCachedArticles()
	}
}

func updateArticleR(articleID uint, upds map[string]interface{}) {
	rdb.HSet(ctx, "articles:"+strconv.Itoa(int(articleID)), upds)
}

func getArticleR(articleID uint) (article Article) {
	res, _ := rdb.HGetAll(ctx, "articles:"+strconv.Itoa(int(articleID))).Result()
	article.ID = articleID
	article.Title = res["title"]
	article.Content = res["content"]
	article.Author = res["author"]
	return
}
func getPopularArticlesR() (articles []Article) {
	articleIDstrs, _ := rdb.SMembers(ctx, "cached_ids").Result()
	for _, articleIDstr := range articleIDstrs {
		articleID, _ := strconv.Atoi(articleIDstr)
		res, _ := rdb.HGetAll(ctx, "articles:"+articleIDstr).Result()
		var article Article
		article.ID = uint(articleID)
		article.Title = res["title"]
		article.Content = res["content"]
		article.Author = res["author"]
		articles = append(articles, article)
	}
	return
}

func isCached(articleID uint) bool {
	isCached, _ := rdb.SIsMember(ctx, "cached_ids", articleID).Result()
	return isCached
}
