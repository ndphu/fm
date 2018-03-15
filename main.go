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

	authorized := r.Group("/pfm/api")
	authorized.Use(filter.AuthFilter())

	// login
	loginService := service.NewLoginService(db)
	loginController := controller.NewLoginController(loginService)
	loginController.Init(r)

	// file
	fileService := service.NewFileService(db)
	fileService.Init()
	controller.NewFileController(fileService, r)

	r.Run()
}
