package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tamabsndra/miniproject/miniproject-backend/services"
	"github.com/tamabsndra/miniproject/miniproject-backend/utils"
)

func AuthMiddleware(jwtSecret string, tokenService *services.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("authToken")
		if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized "})
            c.Abort()
            return
        }

		if tokenService.IsTokenBlacklisted(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token has been revoked"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("token", token)
		c.Next()
	}
}
