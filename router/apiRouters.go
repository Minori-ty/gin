package router

import (
	"fmt"
	"gin/controllers/admin"
	"gin/controllers/api"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	c.Next()

	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func ApiRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", initMiddleware, api.UserController{}.Index)

		r.GET("/article", initMiddleware, api.UserController{}.Article)

		r.POST("/login", initMiddleware, admin.UserController{}.Login)

		r.POST("/upload", api.UserController{}.Upload)
	}
}
