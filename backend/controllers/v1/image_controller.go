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
			"err":     err.Error(),
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
