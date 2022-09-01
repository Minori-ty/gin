package main

import (
	"fmt"
	"gin/common"
	"gin/middlewares"
	"gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(common.Cors())
	r.Use(middlewares.GlobalMiddleware)

	router.ApiRoutersInit(r)

	router.UserRoutersInit(r)

	fmt.Println("http://localhost:4000")
	r.Run(":4000")
}
