package main

import (
	"log"

	"github.com/chyuhung/my-dashboard/api"
	"github.com/chyuhung/my-dashboard/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func main() {
	// 加载配置
	config := config.LoadConfig()

	// 使用配置中的认证参数进行认证
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: config.AuthURL,
		Username:         config.Username,
		Password:         config.Password,
		DomainName:       config.DomainName,
		AllowReauth:      true,
		Scope: &gophercloud.AuthScope{
			ProjectName: config.ProjectName,
			DomainName:  config.DomainName},
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		log.Fatal("Failed to create OpenStack client: ", err)
	}

	// 创建Gin引擎
	router := SetupRouter()

	// 设置路由规则
	router.GET("/volumes", api.ListVolumes)

	// 启动HTTP服务器
	log.Println("Server started on localhost:8080")
	log.Fatal(router.Run(":8080"))
}
