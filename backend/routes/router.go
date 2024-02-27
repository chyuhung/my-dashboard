package routes

import (
	"github.com/chyuhung/my-dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/logout", controllers.Logout)
		authRoutes.POST("/register", controllers.Register)
	}

	// flavor
	flavorRoutes := router.Group("/flavors")
	{
		flavorRoutes.GET("/", controllers.ListFlavors)
		flavorRoutes.GET("/:id", controllers.GetFlavor)
	}

	// instance
	instancesRoutes := router.Group("/instances")
	{
		// 所有虚拟机列表
		instancesRoutes.GET("/", controllers.ListInstances)
		// 查询单个
		instancesRoutes.GET("/:id", controllers.GetInstance)
		// 新增
		instancesRoutes.POST("/", controllers.CreateInstance)
		// 编辑
		instancesRoutes.PUT("/:id", controllers.UpdateInstance)
	}

	// volume
	volumeRoutes := router.Group("/volumes")
	{
		volumeRoutes.GET("/", controllers.ListVolumes)
		volumeRoutes.GET("/:id", controllers.GetVolume)
	}
	// network
	networkRoutes := router.Group("/networks")
	{
		networkRoutes.GET("/", controllers.ListNetworks)
		networkRoutes.GET("/:id", controllers.GetNetwork)
	}
	// image
	imageRoutes := router.Group("/images")
	{
		imageRoutes.GET("/", controllers.ListImages)
		imageRoutes.GET("/:id", controllers.GetImage)
	}
	// 其他路由设置...
}
