package article

import (
	"mime/multipart"
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"
)

type Usecase interface {
	CreateArticle(*model.Article) error
	GetArticles(article *[]model.Article) (err error)
	GetArticleById(uint, *model.Article) (err error)
	UploadImages(article int, files []*multipart.FileHeader) error
}

type service struct {
	articleRepo    repository.ArticleRepository
	articlePersist persistence.ArticlePersistence
}

func Initialize(
	articleRepo repository.ArticleRepository,
	articlePersist persistence.ArticlePersistence,
) Usecase {
	return &service{
		articleRepo,
		articlePersist,
	}
}
