package filter

import (
	//"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCorsFilter(r *gin.Engine) {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length"}
	c.AllowCredentials = true
	r.Use(cors.New(c))
}
