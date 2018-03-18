package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/fm/model"
	"github.com/ndphu/fm/service"
	"gopkg.in/mgo.v2"
	"net/http"
	"strings"
)

var ()

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
	// file
	r.POST("/pfm/api/files", c.createFile)
	r.GET("/pfm/api/file/:id", c.getFileById)
	r.PUT("/pfm/api/file/:id", c.updateFileById)
	r.DELETE("/pfm/api/file/:id", c.deleteFileById)

	// children
	r.GET("/pfm/api/file/:id/children", c.getChildren)
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
		if err = c.validateFileObject(&newFile); err != nil {
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
	g.JSON(500, gin.H{"err": "method not impleted"})
}

func (c *FileController) deleteFileById(g *gin.Context) {
	g.JSON(500, gin.H{"err": "method not impleted"})

}

func (c *FileController) getChildren(g *gin.Context) {
	id := g.Param("id")
	list, err := c.fileService.FindChildren(id)
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
		g.JSON(http.StatusOK, list)
	}

}

func (c *FileController) validateFileObject(f *model.File) error {
	if strings.Trim(f.Name, " ") == "" {
		return errors.New("file name can not empty")
	}

	for _, char := range []string{"\\", "|", "/", "*", "?", "\"", "<", ">"} {
		if strings.Index(f.Name, char) >= 0 {
			return errors.New("file name cannot contains '" + char + "'")
		}
	}

	if f.ParentId.Hex() == "" {
		return errors.New("parentId is required")
	} else {
		parent, err := c.fileService.FileFileById(f.ParentId.Hex())
		if err == mgo.ErrNotFound {
			return errors.New("parent not exists")
		}
		if parent.Type != model.TYPE_FOLDER {
			return errors.New("parent should be a folder")
		}
	}
	return nil
}
