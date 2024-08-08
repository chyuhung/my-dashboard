package v1

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
)

func CreateServerHandler(c *gin.Context) {
	// 绑定json数据到结构体
	var csr models.CreateServerRequest
	if err := c.ShouldBindJSON(csr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无法解析请求数据",
		})
		return
	}
	// 从csr读取创建server的必要信息
	// 检查image
	_, err := services.GetImage(csr.OS)
	if err == nil {
		// image不存在
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "image不存在",
		})
		return
	}
	// 检查flavor
	_, err = services.GetFlavor(csr.Flavor)
	if err == nil {
		// flavor不存在
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "flavor不存在",
		})
		return
	}
	// 检查network
	for _, n := range csr.Networks {
		_, err = services.GetNetwork(n.Vlan)
		if err == nil {
			// network不存在
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "network不存在",
			})
			return
		}
	}
	// volume耗时最长，首先进行创建
	// 添加随机字符串后缀5位，保证名称唯一
	for _, v := range csr.Volumes {
		suffix := GenerateRandomString(5)
		opts := volumes.CreateOpts{Size: v.Size, Name: csr.Hostname + "-datavol-" + suffix}
		volume, err := services.CreateVolume(opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "创建volume失败",
				"err":     err.Error(),
			})
			return
			// 创建volume
		}
		// 从openstack获取创建server的必要信息
		// 调用openstack创建server
		// 返回创建结果

	}
}

// GenerateRandomString 生成一个指定长度的随机字符串
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
