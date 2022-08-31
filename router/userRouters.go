package router

import (
	"gin/controllers/admin"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	{
		userRouters.GET("/", admin.UserController{}.Index)

		userRouters.GET("/editor:id", admin.UserController{}.Article)
	}
}
