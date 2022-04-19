package persistence

import (
	"fmt"
	"twof/blog-api/internal/constant/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ArticlePersistence interface {
	CreateArticle(article *model.Article) (err error)
	GetAllArticles(article *[]model.Article) (err error)
	GetArticleById(articleId uint, article *model.Article) (err error)
	GetArticleByQuery(article *model.Article) (err error)
	UpdateImage(articleId int, images []string) (err error)
}

type articlePersistence struct {
	db *gorm.DB
}

func ArticleInit(db *gorm.DB) ArticlePersistence {
	return &articlePersistence{
		db,
	}
}

func (ar articlePersistence) CreateArticle(article *model.Article) (err error) {

	if err = ar.db.Create(article).Error; err != nil {
		fmt.Println("her is the error")
		return err
	}

	return nil
}

func (ar *articlePersistence) GetArticleByQuery(article *model.Article) (err error) {
	if err = ar.db.Where("title = ?", article.Title).First(article).Error; err != nil {
		return err
	}

	return nil
}

func (ar *articlePersistence) GetArticleById(articleId uint, article *model.Article) (err error) {
	if err = ar.db.Where("ID = ?", articleId).First(article).Error; err != nil {
		return err
	}

	return nil
}

func (ar *articlePersistence) GetAllArticles(article *[]model.Article) (err error) {
	if err = ar.db.Find(article).Error; err != nil {
		return err
	}

	return nil
}

func (ar *articlePersistence) UpdateImage(articleId int, images []string) (err error) {
	var article *model.Article
	id := uint(articleId)
	if err = ar.db.Where("ID = ?", id).First(&article).Update("Image", pq.Array(images)).Error; err != nil {
		return err
	}

	return nil
}
