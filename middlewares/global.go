package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleware(c *gin.Context) {
	fmt.Println("全局中间件")
}
