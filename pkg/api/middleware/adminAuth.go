package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	adminID, role, err := ValidateJWT(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role is not admin%v", "err": err.Error()})
		c.Abort()
		return
	}

	c.Set("adminId", adminID)
	c.Next()
}
