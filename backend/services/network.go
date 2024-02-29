package services

import (
	"fmt"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

var networkClient *gophercloud.ServiceClient

func SetNetworkClient(client *gophercloud.ServiceClient) {
	networkClient = client
}

func GetNetworkId(networkName string) (string, error) {
	listOpts := networks.ListOpts{
		Limit: 99999,
	}
	// 发送查询 network 列表请求
	allPages, err := networks.List(networkClient, listOpts).AllPages()
	if err != nil {
		// 处理查询 network 列表失败的情况
		return "", err
	}
	networks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return "", err
	}
	// 遍历 network 列表，查找 network 名称匹配的 network
	for _, network := range networks {
		if network.Name == networkName {
			// 返回 network ID
			return network.ID, nil
		}
	}
	return "", fmt.Errorf("network %s not found", networkName)
}

func ListNetworks() ([]*models.Network, error) {
	var data []*models.Network
	listOpts := networks.ListOpts{}

	allPages, err := networks.List(networkClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return nil, err
	}
	for _, network := range allNetworks {
		data = append(data, &models.Network{
			Vlan: network.Name,
			Ip:   "",
		})
	}
	return data, nil
}

func GetNetwork(id string) (networks.Network, error) {
	network, err := networks.Get(networkClient, id).Extract()
	if err != nil {
		return networks.Network{}, err
	}
	return *network, nil
}
