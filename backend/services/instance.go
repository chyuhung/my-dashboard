package services

import (
	"github.com/chyuhung/my-dashboard/models"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

var computeClient *gophercloud.ServiceClient

func SetComputeClient(client *gophercloud.ServiceClient) {
	computeClient = client
}

func GetInstances() ([]*models.Instance, error) {
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

func CreateInstance(name string) (*models.Instance, error) {
	// 创建实例请求参数
	createOpts := servers.CreateOpts{
		Name:      name,
		FlavorRef: "flavor_id", // 替换为实际的 flavor ID
		ImageRef:  "image_id",  // 替换为实际的 image ID
		// 其他实例配置参数
	}

	// 发送创建实例请求
	server, err := servers.Create(computeClient, createOpts).Extract()
	if err != nil {
		// 处理创建实例失败的情况
		return nil, err
	}

	// 将响应数据转换为自定义的 Instance 模型
	instance := &models.Instance{
		ID:   server.ID,
		Name: server.Name,
		// 其他实例数据
	}

	return instance, nil
}

func UpdateInstance(id string, name string) (*models.Instance, error) {
	// 更新实例请求参数
	updateOpts := servers.UpdateOpts{
		Name: name,
		// 其他要更新的实例配置参数
	}

	// 发送更新实例请求
	server, err := servers.Update(computeClient, id, updateOpts).Extract()
	if err != nil {
		// 处理更新实例失败的情况
		return nil, err
	}

	// 将响应数据转换为自定义的 Instance 模型
	instance := &models.Instance{
		ID:   server.ID,
		Name: server.Name,
		// 其他实例数据
	}

	return instance, nil
}
