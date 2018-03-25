package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/controller"
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/filter"
	"github.com/ndphu/fm/service"
)

func main() {
	fmt.Println("Starting application...")
	db := dao.NewDB()
	if err := db.Init(); err != nil {
		panic(err)
	}
	defer db.Shutdown()

	r := gin.Default()
	filter.SetupCorsFilter(r)

	authFilter := filter.NewAuthFilter()
	r.Use(authFilter.AuthFilter())

	// token
	tokenService := service.NewTokenService()

	authFilter.SetTokenService(tokenService)

	// user
	userService := service.NewUserService(db)

	// login
	loginService := service.NewLoginService(db)
	loginService.SetUserService(userService)
	loginController := controller.NewLoginController()
	loginController.SetLoginService(loginService)
	loginController.SetTokenService(tokenService)
	loginController.Init(r)

	// file
	fileService := service.NewFileService(db)
	fileService.Init()
	controller.NewFileController(fileService, r)

	// user
	userController := controller.NewUserController()
	userController.SetUserService(userService)
	userController.Init(r)

	r.Run()
}
