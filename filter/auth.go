package filter

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/model"
	"github.com/ndphu/fm/service"
	"strings"
)

type AuthFilter struct {
	tokenService *service.TokenService
}

func NewAuthFilter() *AuthFilter {
	return &AuthFilter{}
}

func (f *AuthFilter) SetTokenService(s *service.TokenService) {
	f.tokenService = s
}

func (f *AuthFilter) AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthFilter: checking path " + c.Request.URL.Path)
		if strings.ToUpper(c.Request.Method) == "OPTIONS" {
			// ignore for CORS
			c.Next()
		} else if isIgnorePath(c.Request.URL.Path) {
			fmt.Println("AuthFilter: ignore path " + c.Request.URL.Path)
			c.Next()
		} else {
			auth := c.GetHeader("Authorization")
			if usr, err := f.validateToken(auth); err == nil {
				c.Set("user", usr)
				c.Next()
			} else {
				c.JSON(401, gin.H{"err": "not authorized. " + err.Error()})
				c.Abort()
			}
		}
	}
}

func isIgnorePath(path string) bool {
	for _, p := range []string{
		"/pfm/api/login/facebook",
		"/pfm/api/admin/login",
	} {
		if path == p {
			return true
		}
	}
	return false
}

func (f *AuthFilter) validateToken(auth string) (*model.User, error) {
	if auth == "" {
		return nil, errors.New("empty header")
	}
	tokenString := strings.Trim(strings.TrimPrefix(auth, "Bearer "), " ")
	return f.tokenService.ValidateToken(tokenString)
}
