package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID      uint   `gorm:"primary key" redis:"id"`
	Title   string `redis:"title"`
	Content string `redis:"content"`
	Author  string `redis:"author"`
	Visited int    `redis:"visited"`
}

var db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db240830_2?charset=utf8mb4&parseTime=True&loc=Local&&timeout=10s"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("连接数据库成功")
	}

	db.AutoMigrate(&Article{})
}

func createArticle(article *Article) {
	db.Create(article)
	createArticleR(article.ID)
}

func getAllArticles() (articles []Article) {
	db.Model(&Article{}).Find(&articles)
	return
}

func getPopularArticles() (articles []Article) {
	articles = getPopularArticlesR()
	return
}

func getArticle(articleID uint) (article Article) {
	visitArticleR(articleID)
	if isCached(articleID) {
		fmt.Println("get article from cache")
		article = getArticleR(articleID)
	} else {
		fmt.Println("get article from db")
		db.First(&article, articleID)
	}
	return
}

func updateArticle(articleID uint, upds map[string]interface{}) {
	db.Model(&Article{}).Where("id = ?", articleID).Updates(upds)
	if isCached(articleID) {
		updateArticleR(articleID, upds)
	}
}

func updateVisited(articleID uint, visited int) {
	db.Model(&Article{}).Where("id = ?", articleID).Update("visited", visited)
}

func deleteArticle(articleID uint) {
	db.Delete(&Article{}, articleID)
	deleteArticleR(articleID)
}

// func visitArticle(times int, articleID uint) {
// 	var visited int
// 	db.Model(&Article{}).Select("visited").Find(&visited, articleID)
// 	db.Model(&Article{}).Where("id = ?", articleID).Update("visited", visited+times)
// }

// func popularArticles(num int) (articles []Article) {
// 	db.Model(&Article{}).Order("visited DESC").Find(&articles).Limit(num)
// 	return
// }
