package middleware

import (
	"daily-worker-roster-management-system/handlers/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "authorization header required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "authorization header format must be Bearer {token}")
			c.Abort()
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid token claims")
			c.Abort()
			return
		}

		u, ok := claims["user"]
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid token user_id")
			c.Abort()
			return
		}
		// Store user ID in context so handlers can use it
		c.Set("user", u)

		c.Next()
	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":          403,
				"status":        "error",
				"error_message": "Forbidden: admin access only",
			})
			return
		}
		roleId := user.(map[string]interface{})["role_id"].(float64)
		if roleId != 1 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":          403,
				"status":        "error",
				"error_message": "Forbidden: admin access only",
			})
		}
		c.Next()
	}
}
