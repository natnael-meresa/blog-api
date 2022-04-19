package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/module/article"

	"github.com/gin-gonic/gin"
)

type ArticleHandler interface {
	Create(*gin.Context)
	GetAll(c *gin.Context)
	UploadImage(c *gin.Context)
}

type articleHandler struct {
	articleCase article.Usecase
}

func ArticleInit(articleCase article.Usecase) ArticleHandler {
	return &articleHandler{
		articleCase,
	}
}
func (ar *articleHandler) Create(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}

	err := ar.articleCase.CreateArticle(&article)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)

		return
	}
	state.ResJsonData(c, "Article Created", http.StatusOK, article)
}

func (ar *articleHandler) GetAll(c *gin.Context) {
	var article []model.Article

	err := ar.articleCase.GetArticles(&article)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
	}

	state.ResJsonData(c, "Article fetched", http.StatusOK, article)
}

func (ar *articleHandler) UploadImage(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}
	files := form.File["files"]

	err = ar.articleCase.UploadImages(articleId, files)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}

	state.ResJsonData(c, fmt.Sprintf("upload %d files", len(files)), http.StatusOK, len(files))

}
