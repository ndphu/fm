package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/model"
	"github.com/ndphu/fm/service"
	"gopkg.in/mgo.v2"
	"net/http"
)

var (
	PATH = "/api/file"
)

type FileController struct {
	fileService *service.FileService
}

func NewFileController(fs *service.FileService, r *gin.Engine) *FileController {
	c := FileController{
		fileService: fs,
	}

	c.initRouter(r)

	return &c
}

func (c *FileController) initRouter(r *gin.Engine) {
	r.POST(PATH+"/", c.createFile)
	r.GET(PATH+"/:id", c.getFileById)
	r.PUT(PATH+"/:id", c.updateFileById)
}

func (c *FileController) getFileById(g *gin.Context) {
	id := g.Param("id")
	f, err := c.fileService.FileFileById(id)
	if err != nil {
		var status int
		if err == mgo.ErrNotFound {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		g.JSON(status, gin.H{
			"err": err.Error(),
		})

	} else {
		g.JSON(http.StatusOK, f)
	}

}

func (c *FileController) createFile(g *gin.Context) {
	var newFile model.File
	if err := g.ShouldBindJSON(&newFile); err == nil {
		if err = validateFileObject(&newFile); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		err = c.fileService.CreateFile(&newFile)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		} else {
			g.JSON(http.StatusCreated, newFile)
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

func (c *FileController) updateFileById(g *gin.Context) {
}

func validateFileObject(f *model.File) error {
	if f.IsRoot {
		return errors.New("you can't create root folder")
	}
	return nil
}
