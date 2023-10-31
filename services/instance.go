package services

import (
	"github.com/chyuhung/my-dashboard/backend/models"
	"github.com/gophercloud/gophercloud"
)

func GetInstances(computeClient *gophercloud.ServiceClient) ([]*models.Instance, error) {
	// 获取实例逻辑
}

func CreateInstance(computeClient *gophercloud.ServiceClient, name string) (*models.Instance, error) {
	// 创建实例逻辑
}

func UpdateInstance(computeClient *gophercloud.ServiceClient, id string, name string) (*models.Instance, error) {
	// 更新实例逻辑
}

// 其他服务函数...
