package main

import (
	"fmt"
	"gin/common"
	"gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(common.Cors())

	router.ApiRoutersInit(r)

	router.UserRoutersInit(r)

	fmt.Println("http://localhost:3000")
	r.Run(":4000")
}
