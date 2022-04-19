package userInit

import (
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/glue/enforcer"
	"twof/blog-api/internal/glue/routing"
	"twof/blog-api/internal/handler/rest"
	"twof/blog-api/internal/module/user"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(en enforcer.CasbinMiddleware, g *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	databaseUser := persistence.UserInit(db)

	repo := repository.UserInit("secret")
	usecase := user.Initialize(en, repo, databaseUser)
	handler := rest.UserInit(usecase)

	routing.UserRouting(g, handler)

}
