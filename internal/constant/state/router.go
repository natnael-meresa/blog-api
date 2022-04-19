package state

import "github.com/gin-gonic/gin"

type Router struct {
	Method     string
	Path       string
	Handler    gin.HandlerFunc
	Middleware []gin.HandlerFunc
}
