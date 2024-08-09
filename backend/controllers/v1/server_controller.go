package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
)

func GetServersHandler(c *gin.Context) {
	// 获取所有server
	servers, err := services.GetServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "无法获取服务器列表",
		})
		return
	}
	c.JSON(http.StatusOK, servers)
}

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
		// 需要动态交互获取创建volume的必要信息
		opts := volumes.CreateOpts{Size: v.Size, Name: csr.Hostname + "-datavol-" + suffix}
		_, err := services.CreateVolume(opts)
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

// SearchServersHandler 通过名称或 IP 地址搜索服务器
func SearchServersHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询参数不能为空"})
		return
	}

	servers, err := services.SearchServers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "搜索服务器失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, servers)
}
