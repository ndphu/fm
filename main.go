package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/controller"
	"github.com/ndphu/fm/dao"
	"github.com/ndphu/fm/service"
)

func main() {
	fmt.Println("Starting application...")
	DAO := dao.NewDB()
	if err := DAO.Init(); err != nil {
		panic(err)
	}
	defer DAO.Shutdown()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fileService := service.NewFileService(DAO)
	fileService.Init()

	controller.NewFileController(fileService, r)

	r.Run()
}
