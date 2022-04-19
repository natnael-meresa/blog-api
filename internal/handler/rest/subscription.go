package rest

import (
	"net/http"
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/module/subscription"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler interface {
	Create(*gin.Context)
	// GetAll(c *gin.Context)
}

type subscriptionHandler struct {
	subscriptionCase subscription.Usecase
}

func SubscriptionInit(subscriptionCase subscription.Usecase) SubscriptionHandler {
	return &subscriptionHandler{
		subscriptionCase,
	}
}
func (su *subscriptionHandler) Create(c *gin.Context) {
	var subscription model.Subscription
	if err := c.ShouldBind(&subscription); err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}

	err := su.subscriptionCase.CreateSubscription(&subscription)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)

		return
	}
	state.ResJsonData(c, "subscription Created", http.StatusOK, subscription)
}

// func (ar *articleHandler) GetAll(c *gin.Context) {
// 	var article []model.Article

// 	err := ar.articleCase.GetArticles(&article)

// 	if err != nil {
// 		state.ResErr(c, err, http.StatusBadRequest)
// 	}

// 	state.ResJsonData(c, "Article fetched", http.StatusOK, article)
// }
