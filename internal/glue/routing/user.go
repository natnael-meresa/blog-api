package routing

import (
	"net/http"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func UserRouting(g *gin.RouterGroup, handler rest.UserHandler) {

	userRoutes := []state.Router{
		{
			Method:     http.MethodPost,
			Path:       "/user/login",
			Handler:    handler.LogUserHandler,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodPost,
			Path:       "/user/register",
			Handler:    handler.RegistrationHandler,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodPost,
			Path:       "/user/token/refresh",
			Handler:    handler.Refresh,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodPut,
			Path:       "user/reset-link",
			Handler:    handler.ResetLink,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method:     http.MethodPut,
			Path:       "user/password-reset",
			Handler:    handler.PasswordReset,
			Middleware: []gin.HandlerFunc{},
		},
		// {
		// 	Method:     http.MethodGet,
		// 	Path:       "/",
		// 	Handler:    handler.GetUsers,
		// 	Middleware: []gin.HandlerFunc{},
		// },
	}

	gUser := g.Group("auth")
	{
		for _, router := range userRoutes {
			router.Middleware = append(router.Middleware, router.Handler)
			gUser.Handle(router.Method, router.Path, router.Middleware...)
		}
	}

}
