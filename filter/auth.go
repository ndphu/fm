package filter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	LOGIN_REDIRECT = "https://www.facebook.com/v2.12/dialog/oauth?client_id=568365100192445&redirect_uri=http://localhost:8080/pfm/login&state={st=state123abc,ds=123456789}&response_type=token"
)

func AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthFilter is running")
		auth := c.GetHeader("Authorization")
		if isAuthValid(auth) {
			c.Next()
		} else {
			c.JSON(401, gin.H{"err": "not authorized"})
		}
	}
}

func isAuthValid(auth string) bool {
	if auth == "" {
		return false
	}

	return true
}
