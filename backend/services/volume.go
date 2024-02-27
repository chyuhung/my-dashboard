package services

import (
	"fmt"
	"time"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
)

var volumeClient *gophercloud.ServiceClient

func SetVolumeClient(client *gophercloud.ServiceClient) {
	volumeClient = client
}

// 查询单个volume信息
func GetVolume(id string) (volumes.Volume, error) {
	volume, err := volumes.Get(volumeClient, id).Extract()
	if err != nil {
		return volumes.Volume{}, err
	}
	return *volume, nil
}

func GetVolumeTypeId(volumeTypeName string) (string, error) {
	listOpts := volumetypes.ListOpts{
		Limit: -1,
	}
	// 发送查询 volumetype 列表请求
	allPages, err := volumetypes.List(volumeClient, listOpts).AllPages()
	if err != nil {
		// 处理查询 volumetype 列表失败的情况
		return "", err
	}
	volumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return "", err
	}
	// 遍历 volumetype 列表，查找 volumetype 名称匹配的 volumetype
	for _, volueType := range volumeTypes {
		if volueType.Name == volumeTypeName {
			// 返回 volumetype ID
			return volueType.ID, nil
		}
	}
	return "", fmt.Errorf("volueType %s not found", volumeTypeName)
}

func ListVolumes() ([]*models.Volume, error) {
	var data []*models.Volume
	listOpts := volumes.ListOpts{}

	allPages, err := volumes.List(volumeClient, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	fmt.Println(allPages)
	allVolumes, err := volumes.ExtractVolumes(allPages)
	if err != nil {
		return nil, err
	}
	fmt.Println(allVolumes)
	for _, volume := range allVolumes {
		data = append(data, &models.Volume{
			Name: volume.Name,
			Size: volume.Size,
			Type: volume.VolumeType,
		})
	}
	return data, nil
}

func CheckVolStatus(id string, ch chan bool) {
	for {
		// 检查 vol 状态是否为 available
		vol, err := GetVolume(id)
		if err != nil {
			ch <- false // 将结果发送到通道，表示失败
			return
		}
		status := vol.Status

		if status == "available" {
			ch <- true // 将结果发送到通道，表示完成
			return
		} else if status == "error" {
			ch <- false // 将结果发送到通道，表示失败
			return
		}

		time.Sleep(5 * time.Second) // 休眠后再次检查
	}
}

func CreateVolume(name string, size int, volumeType string, imageId string) (volumes.Volume, error) {
	createOpts := volumes.CreateOpts{
		Size:       size,
		Name:       name,
		VolumeType: volumeType,
		ImageID:    imageId,
	}
	volume, err := volumes.Create(volumeClient, createOpts).Extract()
	if err != nil {
		return volumes.Volume{}, err
	}
	return *volume, nil
}
