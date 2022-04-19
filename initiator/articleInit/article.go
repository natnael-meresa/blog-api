package articleInit

import (
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/glue/routing"
	"twof/blog-api/internal/handler/rest"
	"twof/blog-api/internal/module/article"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(g *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&model.Article{})
	databaseProfile := persistence.ArticleInit(db)

	repo := repository.ArticleInit()
	usecase := article.Initialize(repo, databaseProfile)
	handler := rest.ArticleInit(usecase)

	routing.ArticleRouting(g, handler)

}
