package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/service"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(ls *service.LoginService) *LoginController {
	return &LoginController{
		loginService: ls,
	}
}

func (ctl *LoginController) Init(r *gin.Engine) {
	lg := r.Group("/pfm/login")
	{
		lg.GET("/callback", ctl.LoginCallbackHandler)
		lg.GET("", ctl.LoginHandler)
	}
}

func (ctl *LoginController) LoginCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	fmt.Println("Got FB code " + code)
	ctl.loginService.ProcessAccessCode(code)

}

func (ctl *LoginController) LoginHandler(c *gin.Context) {
	c.Redirect(301, "https://www.facebook.com/v2.12/dialog/oauth?client_id=568365100192445&redirect_uri=http://localhost:8080/pfm/login/callback&state={st=state123abc,ds=123456789}&response_type=code")
}
