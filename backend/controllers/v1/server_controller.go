package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
)

// 通过主机名或ip查询主机信息，返回自定义model类型数据
func GetServersHandler(c *gin.Context) {
	query := c.Query("query")
	servers, err := services.GetServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "查询主机失败",
			"err":     err,
		})
		return
	}
	// 自定义model类型数据
	var svs []models.CreateServerRequest
	for _, s := range servers {
		if s.Name == query || s.AccessIPv4 == query {
			svs = append(svs, models.CreateServerRequest{
				Hostname: s.Name,
				OS:       s.Image["Name"].(string),
				Flavor:   s.Flavor["Name"].(string),
			})
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询主机成功",
		"data":    svs,
	})
}

// CheckServerHandler 检查服务器
func CheckServerHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询参数不能为空"})
		return
	}

	servers, err := services.GetServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取主机名列表败",
			"error":   err,
		})
		return
	}
	for _, s := range servers {
		if s.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "主机名不可用",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "主机名可用",
	})
}

func CreateHandler(c *gin.Context) {
	// 绑定json数据到结构体
	var req models.CreateServerRequest
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无法解析请求数据",
		})
		return
	}
	// 从req读取创建server的必要信息
	// 检查image
	_, err := services.GetImage(req.OS)
	if err == nil {
		// image不存在
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "image不存在",
		})
		return
	}
	// 检查flavor
	_, err = services.GetFlavor(req.Flavor)
	if err == nil {
		// flavor不存在
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "flavor不存在",
		})
		return
	}
	// 检查network
	for _, n := range req.Networks {
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
	for _, v := range req.Volumes {
		suffix := GenerateRandomString(5)
		// 需要动态交互获取创建volume的必要信息
		opts := volumes.CreateOpts{Size: v.Size, Name: req.Hostname + "-datavol-" + suffix}
		_, err := services.CreateVolume(opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "创建volume失败",
				"err":     err,
			})
			return
		}
		// 创建volume
		// 从openstack获取创建server的必要信息
		// 调用openstack创建server
		// 返回创建结果

	}
}
