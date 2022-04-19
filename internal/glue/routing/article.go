package routing

import (
	"net/http"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func ArticleRouting(g *gin.RouterGroup, handler rest.ArticleHandler) {

	articleRoutes := []state.Router{
		{
			Method:     http.MethodPost,
			Path:       "/",
			Handler:    handler.Create,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodGet,
			Path:       "/",
			Handler:    handler.GetAll,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodPut,
			Path:       "/:id/image",
			Handler:    handler.UploadImage,
			Middleware: []gin.HandlerFunc{},
		},
	}

	gArticle := g.Group("articles")
	{
		for _, router := range articleRoutes {
			router.Middleware = append(router.Middleware, router.Handler)
			gArticle.Handle(router.Method, router.Path, router.Middleware...)
		}
	}

}
