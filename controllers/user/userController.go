package user

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	userList := models.User{}
	models.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"resule": userList,
	})
}
