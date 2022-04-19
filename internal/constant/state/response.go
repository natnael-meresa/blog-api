package state

import (
	"github.com/gin-gonic/gin"
)

func ResErr(c *gin.Context, err error, status int) {
	c.JSON(status, err.Error())

	c.Abort()
}

func ResJson(c *gin.Context, msg string, status int) {
	c.JSON(status, msg)
	c.Abort()
}

// Response json data with status code
func ResJsonData(c *gin.Context, msg string, status int, v interface{}) {
	c.JSON(status, gin.H{"message": msg, "data": v})
	c.Abort()
}
