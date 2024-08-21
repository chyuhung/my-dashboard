package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
)

func GetServersHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询参数不能为空"})
		return
	}

	servers, err := services.SearchServers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "搜索实例失败",
			"error":   err.Error(),
		})
		return
	}
	// TODO: 转换成model.Server
	var svs []models.CreateServerRequest
	for _, s := range servers {
		svs = append(svs, models.CreateServerRequest{
			Hostname: s.Name,
			//Flavor:   s.Flavor,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "搜索成功",
		"data":    svs,
	})

}

func CreateHandler(c *gin.Context) {
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

// CheckServerHandler 检查服务器
func CheckServerHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询参数不能为空"})
		return
	}

	servers, err := services.SearchServers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "搜索实例失败",
			"error":   err.Error(),
		})
		return
	}
	for _, s := range servers {
		if s.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "名称不可用",
			})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "名称可用",
	})
}
