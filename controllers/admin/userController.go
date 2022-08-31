package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	id := c.DefaultQuery("id", "1")
	c.JSON(200, gin.H{
		"id":  id,
		"msg": "success",
	})
}

type page struct {
	Title string `json:"title"`
	Age   int    `json:"age"`
}

func (con UserController) Article(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":  id,
		"msg": "success",
	})
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (con UserController) Login(c *gin.Context) {
	user := &UserInfo{}
	// 如果等于nil则是成功
	if err := c.ShouldBind(&user); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  err.Error(),
		})
	}
}
