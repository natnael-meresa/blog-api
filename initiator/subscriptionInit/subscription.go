package subscriptionInit

import (
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/glue/routing"
	"twof/blog-api/internal/handler/rest"
	"twof/blog-api/internal/module/subscription"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(g *gin.RouterGroup, db *gorm.DB) {
	db.AutoMigrate(&model.Subscription{})
	databaseSubscription := persistence.SubscriptionInit(db)

	repo := repository.SubscriptionInit()
	usecase := subscription.Initialize(repo, databaseSubscription)
	handler := rest.SubscriptionInit(usecase)

	routing.SubscriptionRouting(g, handler)

}
