package services

import (
	"github.com/chyuhung/my-dashboard/backend/models"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

var computeClient *gophercloud.ServiceClient

func SetComputeClient(client *gophercloud.ServiceClient) {
	computeClient = client
}

func GetInstances(computeClient *gophercloud.ServiceClient) ([]*models.Instance, error) {
	listOpts := servers.ListOpts{}
	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		return nil, err
	}

	instances := make([]*models.Instance, len(allInstances))
	for i, instance := range allInstances {
		// 将实例数据转换为自定义的 Instance 模型
		instances[i] = &models.Instance{
			// ImageName:      instance.Image.Name,
			// FlavorName:     instance.Flavor.Name,
			Volumes:        map[string]int{"volume1": 50, "volume2": 100}, // 替换为实际的卷数据
			VmName:         instance.Name,
			Networks:       map[string]string{"network1": "subnet1", "network2": "subnet2"}, // 替换为实际的网络数据
			HostName:       instance.Metadata["host_name"],
			VolumeTypeName: instance.Metadata["volume_type_name"],
		}
	}

	return instances, nil
}

func CreateInstance(computeClient *gophercloud.ServiceClient, name string) (*models.Instance, error) {
	// 创建实例逻辑
	return nil, nil
}

func UpdateInstance(computeClient *gophercloud.ServiceClient, id string, name string) (*models.Instance, error) {
	// 更新实例逻辑
	return nil, nil
}

// 其他服务函数...
