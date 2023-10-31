package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/logout", controllers.Logout)
		authRoutes.POST("/register", controllers.Register)
	}
	instancesRoutes := router.Group("/instances")
	{
		instancesRoutes.GET("/", controllers.ListInstances)
		instancesRoutes.POST("/", controllers.CreateInstance)
		instancesRoutes.PUT("/:id", controllers.UpdateInstance)
	}
	// 其他路由设置...
}
