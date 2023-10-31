// api/volumes.go

package api

import (
	"log"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

func ListVolumes(provider *gophercloud.ProviderClient, region string) ([]volumes.Volume, error) {
	var data []volumes.Volume
	// 创建Block Storage服务客户端
	blockstorageClient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: region,
	})
	if err != nil {
		log.Printf("Unable to create Block Storage client: %v", err)
		return data, err
	}

	// 调用Block Storage API获取卷列表
	volumesList, err := volumes.List(blockstorageClient, volumes.ListOpts{}).AllPages()
	if err != nil {
		log.Printf("Unable to list volumes: %v", err)
		return data, err
	}

	data, err = volumes.ExtractVolumes(volumesList)
	if err != nil {
		log.Printf("Unable to extract volumes: %v", err)
		return data, err
	}
	return data, nil
}
