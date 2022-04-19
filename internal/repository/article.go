package repository

import (
	"fmt"
	"twof/blog-api/internal/constant/model"

	"github.com/go-playground/validator/v10"
)

type ArticleRepository interface {
	ValidateArticle(*model.Article) error
}

type articleRepository struct {
}

func ArticleInit() ArticleRepository {
	return &articleRepository{}
}

func (ar *articleRepository) ValidateArticle(article *model.Article) (err error) {
	validate = validator.New()

	err = validate.Struct(article)

	if err != nil {
		fmt.Println(err)
		var Msg string
		for _, err := range err.(validator.ValidationErrors) {
			Msg += err.Field() + " is " + err.Tag() + "\n"
		}
		return fmt.Errorf(Msg)
	}

	return nil
}
