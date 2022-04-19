package routing

import (
	"net/http"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func SubscriptionRouting(g *gin.RouterGroup, handler rest.SubscriptionHandler) {

	articleRoutes := []state.Router{
		{
			Method:     http.MethodPost,
			Path:       "/",
			Handler:    handler.Create,
			Middleware: []gin.HandlerFunc{},
		},
	}

	gArticle := g.Group("subscriptions")
	{
		for _, router := range articleRoutes {
			router.Middleware = append(router.Middleware, router.Handler)
			gArticle.Handle(router.Method, router.Path, router.Middleware...)
		}
	}

}
