package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/model"
	"github.com/ndphu/fm/service"
)

type LoginController struct {
	loginService *service.LoginService
	tokenService *service.TokenService
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (c *LoginController) SetLoginService(s *service.LoginService) {
	c.loginService = s
}

func (c *LoginController) SetTokenService(s *service.TokenService) {
	c.tokenService = s
}

func (ctl *LoginController) Init(r *gin.Engine) {
	lg := r.Group("/pfm/api/login")
	{
		lg.GET("/facebook", ctl.FacebookLoginHandler)

	}
	r.POST("/pfm/api/admin/login", ctl.AdminLoginHandler)
}

func (ctl *LoginController) LoginCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	ctl.loginService.ProcessAccessCode(code)
}

func (ctl *LoginController) FacebookLoginHandler(c *gin.Context) {
	accessToken := c.Query("access_token")
	user := ctl.loginService.ProcessAccessToken(accessToken)
	tokenString, err := ctl.tokenService.CreateToken(user)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"token": tokenString})
}

func (ctl *LoginController) AdminLoginHandler(c *gin.Context) {
	ai := AdminLoginInfo{}
	err := c.ShouldBind(&ai)

	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if ai.User == "admin" && ai.Password == "admin" {
		token, err := ctl.tokenService.CreateToken(&model.User{
			FirstName:  "Admin",
			LastName:   "Admin",
			ExternalId: "",
		})
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{"token": token})
	} else {
		c.JSON(401, gin.H{"err": "Invalid username or password"})
	}
}

type AdminLoginInfo struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
