package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
)

// func GetInstance(c *gin.Context) {
// 	str := c.Param("id")
// 	var data *models.Instance
// 	var err error

// 	// 判断str是否是IP
// 	ip := net.ParseIP(str)
// 	if ip != nil {
// 		data, err = services.GetInstanceByIp(str)
// 	} else {
// 		data, err = services.GetInstanceByName(str)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data":    data,
// 		"message": err,
// 	})
// }

// func ListInstances(c *gin.Context) {
// 	// 获取实例数据
// 	instances, err := services.GetInstances()
// 	if err != nil {
// 		// 处理获取实例数据失败的情况
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "获取实例数据失败",
// 		})
// 		return
// 	}

// 	// 返回实例数据
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": instances,
// 	})
// }

// func CreateInstance(c *gin.Context) {
// 	// 解析请求中的实例数据
// 	var instance models.Instance
// 	if err := c.ShouldBindJSON(&instance); err != nil {
// 		// 处理请求数据解析失败的情况
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "无法解析请求数据",
// 		})
// 		return
// 	}
// 	// TODO: 执行创建实例操作，例如将实例数据保存到数据库
// 	err := services.CreateInstance(&instance)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "创建实例失败",
// 			"err":     err.Error(),
// 		})
// 		return
// 	}
// 	// 创建成功
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "实例创建成功",
// 	})
// }

// func UpdateInstance(c *gin.Context) {
// 	// 获取要更新的实例ID
// 	_ = c.Param("id")

// 	// 解析请求中的实例数据
// 	var updatedInstance models.Instance
// 	if err := c.ShouldBindJSON(&updatedInstance); err != nil {
// 		// 处理请求数据解析失败的情况
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "无法解析请求数据",
// 		})
// 		return
// 	}

// 	// TODO: 执行更新实例操作，例如根据实例ID更新数据库中的实例数据

// 	// 更新成功
// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  "实例更新成功",
// 		"instance": updatedInstance,
// 	})
// }

func CreateServerFromVolume(c *gin.Context) {
	// 解析请求中的实例数据
	var customCreateOpts models.CustomCreateOpts
	if err := c.ShouldBindJSON(customCreateOpts); err != nil {
		// 处理请求数据解析失败的情况
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无法解析请求数据",
		})
		return
	}
	var volumeList []*volumes.Volume
	var blockDevices []*bootfromvolume.BlockDevice
	// 首先执行创建volume，将volume创建成功后的信息组合到customCreateOpts
	for _, customVolumeCreateOpts := range customCreateOpts.VolumeCreateOpts {
		volume, err := services.VolumeCreate(customVolumeCreateOpts.CreateOpts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "创建volume失败",
				"err":     err.Error(),
			})
			return
		}
		volumeList = append(volumeList, volume)
		// 将volumeList改造成blockDevices
		for _, volume := range volumeList {
		}
	}
}
