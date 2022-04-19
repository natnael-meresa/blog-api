package auth

import (
	"fmt"
	"net/http"
	"twof/blog-api/internal/glue"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT(skippers ...glue.SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if glue.SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		const BearerSchema string = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found",
			})
		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token",
			})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Not Authorized",
				})
			} else {
				if token.Valid {
					c.Set("userID", claims["userID"])
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Not Authorized",
					})
				}
			}
		}
	}
}

func ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil

	})
}
