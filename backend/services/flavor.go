package services

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
)

func GetFlavors() ([]flavors.Flavor, error) {
	listOpts := flavors.ListOpts{
		Limit: 99999,
	}
	allPages, err := flavors.ListDetail(computeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	flavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}
	return flavors, nil
}

func GetFlavorId(flavorName string) (string, error) {
	listOpts := flavors.ListOpts{
		Limit: 99999,
	}
	// 发送查询 flavor 列表请求
	allPages, err := flavors.ListDetail(computeClient, listOpts).AllPages()
	if err != nil {
		// 处理查询 flavor 列表失败的情况
		return "", err
	}
	flavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return "", err
	}
	// 遍历 flavor 列表，查找 flavor 名称匹配的 flavor
	for _, flavor := range flavors {
		if flavor.Name == flavorName {
			// 返回 flavor ID
			return flavor.ID, nil
		}
	}
	return "", fmt.Errorf("flavor %s not found", flavorName)
}
