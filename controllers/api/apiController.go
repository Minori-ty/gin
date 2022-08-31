package api

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {

	id := c.DefaultQuery("id", "1")
	c.JSON(200, gin.H{
		"id":  id,
		"msg": "success",
	})
}

type article struct {
	Title string `json:"title"`
	Age   int    `json:"age"`
}

func (con UserController) Article(c *gin.Context) {
	a := article{
		Title: "标题",
		Age:   27,
	}
	c.JSON(200, a)
}
