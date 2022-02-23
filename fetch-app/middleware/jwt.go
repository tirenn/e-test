package middleware

import (
	"net/http"
	"tirenn/test-efishery/fetch-app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Header Not Found"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := utils.JWTAuthService().ValidateToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Invalid"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Invalid"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		role := utils.InterfaceToString(claims["role"])
		if len(roles) > 0 && !isRoles(role, roles) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Role Invalid"})
			return
		}
	}
}

func isRoles(role string, roles []string) bool {
	for _, v := range roles {
		if v == role {
			return true
			break
		}
	}

	return false
}
