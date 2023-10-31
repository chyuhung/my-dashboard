// router.go

package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 注册静态文件目录
	router.Static("/static", "./static")

	return router
}
