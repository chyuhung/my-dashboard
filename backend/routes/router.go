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
		// flavor
		// 所有规格列表
		auth.GET("flavors", v1.ListFlavors)
		// 查询单个规格
		auth.GET("flavor/:id", v1.GetFlavor)

		// instance
		// 所有虚拟机列表
		auth.GET("instances", v1.ListInstances)
		// 查询单个虚拟机
		auth.GET("instance/:id", v1.GetInstance)
		// 新增
		auth.POST("instance/add", v1.CreateInstance)
		// 编辑
		auth.PUT("/:id", v1.UpdateInstance)

		// volume
		auth.GET("volumes", v1.ListVolumes)
		auth.GET("volume/:id", v1.GetVolume)

		// network
		auth.GET("networks", v1.ListNetworks)
		auth.GET("network/:id", v1.GetNetwork)

		// image
		auth.GET("images", v1.ListImages)
		auth.GET("image/:id", v1.GetImage)
	}

	// 无需认证
	r := router.Group("/v1")
	// 登录注册模块
	{
		r.POST("login", v1.Login)
	}

	router.Run(":8080")
}
