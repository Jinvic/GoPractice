package article

import (
	"blog-service/pkg/db"
	"blog-service/pkg/define"
	"blog-service/pkg/logger"
	"blog-service/pkg/models"

	"go.uber.org/zap"
)

func Create(uid uint, req *define.ArticleCreateReq) (uint, error) {
	logger.Logger.Info("Create article", zap.Any("user_id", uid), zap.Any("article", req))
	article := models.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: uid,
	}
	err := db.DB.Create(&article).Error
	if err != nil {
		return 0, err
	}
	return article.ID, nil
}

func View(id uint) (*define.ArticleInfo, error) {
	logger.Logger.Info("View article", zap.Any("article_id", id))
	article := models.Article{}
	err := db.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = db.DB.Where("id = ?", article.AuthorID).First(&user).Error
	if err != nil {
		return nil, err
	}
	articleInfo := define.ArticleInfo{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Author:  user.Username,
	}
	return &articleInfo, nil
}

func Edit(aid uint, req *define.ArticleEditReq) error {
	logger.Logger.Info("Edit article", zap.Any("article_id", aid), zap.Any("article", req))
	article := models.Article{}
	article.ID = aid
	article.Title = req.Title
	article.Content = req.Content
	err := db.DB.Updates(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func Delete(id uint) error {
	logger.Logger.Info("Delete article", zap.Any("article_id", id))
	article := models.Article{}
	article.ID = id
	err := db.DB.Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func List(uid uint, username string) (*[]define.ArticleInfo, error) {
	logger.Logger.Info("List user articles", zap.Any("user_id", uid), zap.Any("username", username))
	articles := []models.Article{}
	err := db.DB.Where("author_id = ?", uid).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	articleInfos := []define.ArticleInfo{}
	for _, article := range articles {
		articleInfo := define.ArticleInfo{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Author:  username,
		}
		articleInfos = append(articleInfos, articleInfo)
	}
	return &articleInfos, nil
}

func ListAll() (*[]define.ArticleInfo, error) {
	logger.Logger.Info("List all articles")
	users := []models.User{}
	err := db.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	username := map[uint]string{}
	for _, user := range users {
		username[user.ID] = user.Username
	}

	articles := []models.Article{}
	err = db.DB.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	articleInfos := []define.ArticleInfo{}
	for _, article := range articles {
		articleInfo := define.ArticleInfo{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Author:  username[article.AuthorID],
		}
		articleInfos = append(articleInfos, articleInfo)
	}
	return &articleInfos, nil
}

func GetAuthorID(aid uint) (uint, error) {
	logger.Logger.Info("Get author id", zap.Any("article_id", aid))
	article := models.Article{}
	article.ID = aid
	err := db.DB.First(&article).Error
	if err != nil {
		return 0, err
	}
	return article.AuthorID, nil
}
