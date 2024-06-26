package services

import (
	"github.com/chyuhung/my-dashboard/models"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

var computeClient *gophercloud.ServiceClient

func SetComputeClient(client *gophercloud.ServiceClient) {
	computeClient = client
}

// 通过ip获取实例信息
func GetInstanceByIp(ip string) (*models.Instance, error) {
	var listOpts = servers.ListOpts{
		IP: ip,
	}
	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		return nil, err
	}
	for _, instance := range allInstances {
		for _, v := range instance.Addresses {
			address, _ := v.(string)
			if ip == address {
				return GetInstance(instance.ID)
			}
		}
	}
	return nil, nil
}

// 通过name获取实例信息
func GetInstanceByName(name string) (*models.Instance, error) {
	var listOpts = servers.ListOpts{
		Name: name,
	}
	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		return nil, err
	}
	for _, instance := range allInstances {
		if name == instance.Name {
			return GetInstance(instance.ID)
		}
	}
	return nil, nil
}

// 获取单个实例信息
func GetInstance(id string) (*models.Instance, error) {
	// 根据id获取实例信息
	server, err := servers.Get(computeClient, id).Extract()
	if err != nil {
		return nil, err
	}

	// 根据models中的instance定义重新组合实例信息
	var result = &models.Instance{}
	// 名称
	result.Name = server.Name
	// 规格
	result.Flavor, _ = server.Flavor["name"].(string)
	// volumes和镜像
	var volumes []*models.Volume
	for v := range server.AttachedVolumes {
		info, _ := GetVolume(server.AttachedVolumes[v].ID)
		_ = append(volumes, info)
	}
	// networks
	// 需优化通过ip查询vlan名称
	var networks []*models.Network
	for _, v := range server.Addresses {
		ip, _ := v.(string)
		_ = append(networks, &models.Network{Vlan: "", Ip: ip})
	}
	// 宿主机
	// 需优化通过id查询宿主机名称
	result.Host = server.HostID
	return result, nil
}

// 获取实例列表
func GetInstances() ([]*models.Instance, error) {
	listOpts := servers.ListOpts{
		Limit: 99999,
	}
	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		return nil, err
	}
	// 存储实例信息
	instances := make([]*models.Instance, len(allInstances))
	for i, instance := range allInstances {
		// flavor name
		flavor, _ := instance.Flavor["original_name"].(string)
		// volume
		volumes := []models.Volume{}
		for _, v := range instance.AttachedVolumes {
			volume, err := GetVolume(v.ID)
			if err != nil {
				break
			}
			volumes = append(volumes, models.Volume{
				Name:  volume.Name,
				Size:  volume.Size,
				Type:  volume.Type,
				Image: volume.Image,
			})
		}
		// network
		networks := []models.Network{}
		for _, v := range instance.Addresses {
			ip, _ := v.(string)
			networks = append(networks, models.Network{Vlan: "", Ip: ip})
		}
		// 将实例数据转换为自定义的 Instance 模型
		instances[i] = &models.Instance{
			Name:     instance.Name,
			Flavor:   flavor,
			Volumes:  volumes,
			Networks: networks,
			Host:     instance.HostID,
		}
	}
	return instances, nil
}

func CreateInstance(instance *models.Instance) error {
	// 获取 flavorid
	flavorId, err := GetFlavorId(instance.Flavor)
	if err != nil {
		return err
	}
	// 获取网络
	var networks []servers.Network
	for _, n := range instance.Networks {
		networkId, err := GetNetworkId(n.Vlan)
		if err != nil {
			return err
		} else {
			networks = append(networks, servers.Network{UUID: networkId, FixedIP: n.Ip})
		}
	}

	// 从系统卷和数据卷创建实例
	// 创建实例请求参数
	createOpts := servers.CreateOpts{
		Name:      instance.Name,
		FlavorRef: flavorId,
		Networks:  networks,
		// 其他实例配置参数
		AvailabilityZone: "nova",
	}

	// 遍历创建volume
	blockDevices := []bootfromvolume.BlockDevice{}
	for i, v := range instance.Volumes {
		// 创建系统卷
		if i == 0 {
			sys, err := CreateVolume(v.Name, v.Size, v.Type, v.Image.Id)
			if err != nil {
				return err
			}
			blockDevices = append(blockDevices, bootfromvolume.BlockDevice{
				BootIndex:  i,
				UUID:       sys.ID,
				SourceType: "volume",
			})
		} else {
			// 创建数据卷
			data, err := CreateVolume(v.Name, v.Size, v.Type, "")
			if err != nil {
				return err
			}
			blockDevices = append(blockDevices, bootfromvolume.BlockDevice{
				BootIndex:  i,
				UUID:       data.ID,
				SourceType: "volume",
			})
		}
	}
	// 检查volume状态，状态available后开始创建虚拟机
	// 创建一组 chan 作为标记
	channels := make([]chan bool, len(blockDevices))

	// 启动协程并将标记通道传递给每个协程
	for i, v := range blockDevices {
		channels[i] = make(chan bool)
		go CheckVolStatus(v.UUID, channels[i])
	}

	// 等待所有标记完成，所有volume完成创建并可用，继续执行
	for _, ch := range channels {
		<-ch // 从通道接收结果
	}

	// 从volume启动实例
	bootFromVolumeExt := bootfromvolume.CreateOptsExt{
		CreateOptsBuilder: createOpts,
		BlockDevice:       blockDevices,
	}

	// 发送创建实例请求
	_, err = bootfromvolume.Create(computeClient, bootFromVolumeExt).Extract()
	if err != nil {
		// 处理创建实例失败的情况
		return err
	}
	return nil
}

func UpdateInstance(id string, name string) error {
	// 更新实例请求参数
	updateOpts := servers.UpdateOpts{
		Name: name,
		// 其他要更新的实例配置参数
	}

	// 发送更新实例请求
	_, err := servers.Update(computeClient, id, updateOpts).Extract()
	if err != nil {
		// 处理更新实例失败的情况
		return err
	}

	return nil
}
