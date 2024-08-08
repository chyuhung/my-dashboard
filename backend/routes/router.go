package routes

import (
	v1 "github.com/chyuhung/my-dashboard/controllers/v1"
	middlewares "github.com/chyuhung/my-dashboard/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() {
	router := gin.Default()
	router.Use(middlewares.Cors())

	// 需要认证
	auth := router.Group("/v1")
	auth.Use(middlewares.JwtToken())
	{
		auth.POST("/servers", v1.CreateServerHandler) // 创建 Server
	}

	// 无需认证的路由
	r := router.Group("/v1")
	// 登录注册模块
	{
		r.POST("login", v1.Login)
	}

	router.Run(":8080")
}
