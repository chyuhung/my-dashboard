package routes

import (
	v1 "github.com/chyuhung/my-dashboard/controllers/v1"
	middlewares "github.com/chyuhung/my-dashboard/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() {
	router := gin.Default()
	router.Use(middlewares.Cors())

	auth := router.Group("/v1")
	// 登录注册模块
	{
		auth.POST("login", v1.Login)
		auth.POST("logout", v1.Logout)
		auth.POST("register", v1.Register)
	}
	r := router.Group("/v1")
	{
		// flavor
		// 所有规格列表
		r.GET("flavors", v1.ListFlavors)
		// 查询单个规格
		r.GET("flavor/:id", v1.GetFlavor)

		// instance
		// 所有虚拟机列表
		r.GET("instances", v1.ListInstances)
		// 查询单个虚拟机
		r.GET("instance/:id", v1.GetInstance)
		// 新增
		r.POST("instance/add", v1.CreateInstance)
		// 编辑
		r.PUT("/:id", v1.UpdateInstance)

		// volume
		r.GET("volumes", v1.ListVolumes)
		r.GET("volume/:id", v1.GetVolume)

		// network
		r.GET("networks", v1.ListNetworks)
		r.GET("network/:id", v1.GetNetwork)

		// image
		r.GET("images", v1.ListImages)
		r.GET("image/:id", v1.GetImage)
	}

	router.Run(":8080")
}
