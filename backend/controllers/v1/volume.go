package controllers

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

// 	// 创建卷
// 	_, err := services.CreateVolume(v.Name, v.Size, v.Type, v.Image.Id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{})
// 	}
// }
