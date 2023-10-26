package main

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

func main() {
	// 加载配置
	config := LoadConfig()

	// 使用配置中的认证参数进行认证
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: config.AuthURL,
		Username:         config.Username,
		Password:         config.Password,
		TenantID:         config.ProjectID,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		log.Fatal("Failed to create OpenStack client: ", err)
	}

	// 创建Block Storage服务客户端
	blockstorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: config.Region,
	})
	if err != nil {
		log.Fatal("Failed to create Block Storage client: ", err)
	}

	// 使用Block Storage客户端调用API
	volumesList, err := volumes.List(blockstorageClient, volumes.ListOpts{}).AllPages()
	if err != nil {
		log.Fatal("Failed to list volumes: ", err)
	}

	volumes, err := volumes.ExtractVolumes(volumesList)
	if err != nil {
		log.Fatal("Failed to extract volumes: ", err)
	}

	log.Printf("Volumes: %+v\n", volumes)
}
