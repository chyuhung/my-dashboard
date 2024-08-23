package v1

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/services"
	"github.com/gin-gonic/gin"
)

func GetImagesHandler(c *gin.Context) {
	imgs, err := services.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取image失败",
			"err":     err,
		})
		return
	}
	var images []string
	for _, img := range imgs {
		images = append(images, img.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取image成功",
		"data":    images,
	})
}

func CheckImageHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询参数不能为空",
		})
		return
	}
	imgs, err := services.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取镜像列表失败",
			"err":     err,
		})
		return
	}
	for _, img := range imgs {
		if img.Name == query {
			c.JSON(http.StatusOK, gin.H{
				"message": "镜像可用",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "镜像不可用",
	})
}
