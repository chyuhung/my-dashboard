package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
)

// func GetVolume(c *gin.Context) {
// 	volumeID := c.Param("volumeID")
// 	volume, err := services.GetVolume(volumeID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "获取卷数据失败",
// 		})
// 		return
// 	}

// 	// 返回数据
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": volume,
// 	})
// }
// func ListVolumes(c *gin.Context) {
// 	volumes, err := services.ListVolumes()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "获取卷数据失败",
// 		})
// 		return
// 	}

// 	// 返回数据
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": volumes,
// 	})
// }

// func CreateVolume(c *gin.Context) {
// 	// 解析请求参数
// 	var v models.Volume
// 	if err := c.ShouldBindJSON(&v); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "请求参数错误",
// 		})
// 		return
// 	}

//		// 创建卷
//		_, err := services.CreateVolume(v.Name, v.Size, v.Type, v.Image.Id)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{})
//		}
//	}
func Create(c *gin.Context) {
	var createOpts volumes.CreateOpts
	if err := c.ShouldBindJSON(&createOpts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无法解析请求数据",
		})
		return
	}
	// 执行创建volume操作
	volume, err := services.VolumeCreate(createOpts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败",
			"err":     err.Error(),
		})
		return
	}
	// 创建成功
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"data":    volume,
	})
}
