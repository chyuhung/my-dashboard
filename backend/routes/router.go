package routes

import (
	v1 "github.com/chyuhung/my-dashboard/controllers/v1"
	middlewares "github.com/chyuhung/my-dashboard/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() {
	router := gin.Default()
	router.Use(middlewares.Cors())

	// 需要认证的路由组
	auth := router.Group("/v1")
	auth.Use(middlewares.JwtToken())
	{
		// 创建实例
		auth.POST("/create", v1.CreateHandler) // 创建 Server
		// 获取信息
		auth.GET("/images", v1.GetImagesHandler)           // 获取镜像
		auth.GET("/flavors", v1.GetFlavorsHandler)         // 获取规格
		auth.GET("/volumetypes", v1.GetVolumeTypesHandler) // 获取卷类型
		auth.GET("/networks", v1.GetNetworksHandler)       // 获取网络
		auth.GET("/volumes", v1.GetVolumesHandler)         // 获取卷列表
		// 检查
		auth.GET("/check/server", v1.CheckServerHandler)         // 获取实例详情
		auth.GET("/check/volumetype", v1.CheckVolumeTypeHandler) // 获取卷类型详情
		auth.GET("/check/flavor", v1.CheckFlavorHandler)         // 获取规格详情
		auth.GET("/check/network", v1.CheckNetworkHandler)       // 获取网络详情
		auth.GET("/check/image", v1.CheckImageHandler)           // 获取镜像详情
		// 指定查询
		auth.GET("/servers", v1.GetServersHandler) // 获取实例（部分）
	}

	// 无需认证的路由
	r := router.Group("/v1")
	{
		r.POST("/login", v1.Login) // 登录
	}

	router.Run(":8080")
}
