package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/model"
	"github.com/ndphu/fm/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctl *UserController) SetUserService(s *service.UserService) {
	ctl.userService = s
}

func (ctl *UserController) Init(r *gin.Engine) {
	g := r.Group("/pfm/api/user")
	{
		g.GET("/current", ctl.GetCurrentUser)
	}
	r.GET("/pfm/api/admin/users", ctl.GetAllUsers)
}

func (ctl *UserController) GetCurrentUser(c *gin.Context) {
	usr, _ := c.Get("user")
	u, err := ctl.userService.FindUserByExternalId(usr.(*model.User).ExternalId)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	} else {
		c.JSON(200, u)
	}
}

func (ctl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctl.userService.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
	} else {
		c.JSON(200, users)
	}
}
