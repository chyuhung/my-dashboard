package services

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
)

var imageClient *gophercloud.ServiceClient

func SetImageClient(client *gophercloud.ServiceClient) {
	imageClient = client
}

func GetImageId(imageName string) (string, error) {
	listOpts := images.ListOpts{
		Limit: -1,
	}
	// 发送查询 image 列表请求
	allPages, err := images.ListDetail(imageClient, listOpts).AllPages()
	if err != nil {
		// 处理查询 image 列表失败的情况
		return "", err
	}
	images, err := images.ExtractImages(allPages)
	if err != nil {
		return "", err
	}
	// 遍历 image 列表，查找 image 名称匹配的 image
	for _, image := range images {
		if image.Name == imageName {
			// 返回 image ID
			return image.ID, nil
		}
	}
	return "", fmt.Errorf("image %s not found", imageName)
}
