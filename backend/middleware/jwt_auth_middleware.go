package middleware

import (
	"net/http"
	"strings"

	"github.com/SanjaySinghRajpoot/backend/models"
	tokenutil "github.com/SanjaySinghRajpoot/backend/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, _ := c.Request.Cookie("token")

		authHeader := cookie.Value
		t := strings.Split(authHeader, ".")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
