package auth

import (
	"net/http"
	"strconv"

	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/glue"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
)

func Authorize(enforcer *casbin.Enforcer, skippers ...glue.SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		if glue.SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userid, existed := c.Get("userID")
		userId := userid.(float64)

		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"msg": "User hasn't logged in yet"})
		}

		// userID := contexts.FromUserID(c.Request.Context())

		err := enforcer.LoadPolicy()

		if err != nil {
			state.ResJson(c, "Failed to load policy from DB", http.StatusServiceUnavailable)

			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method

		if ok, err := enforcer.Enforce(strconv.FormatUint(uint64(userId)+1, 10), p, m); err != nil {
			state.ResJson(c, "Error occured when authorizing user", http.StatusServiceUnavailable)

			return
		} else if !ok {
			state.ResJson(c, "You are not authorized", http.StatusUnauthorized)
			return
		}
		c.Next()

	}
}
