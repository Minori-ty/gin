package router

import (
	"gin/controllers/admin"
	"gin/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	{
		userRouters.GET("/", user.UserController{}.Index)

		userRouters.GET("/editor:id", admin.UserController{}.Article)
	}
}
