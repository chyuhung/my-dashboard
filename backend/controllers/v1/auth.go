package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/middlewares"
	"github.com/chyuhung/my-dashboard/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var token string

	c.ShouldBindJSON(&user)
	// 获取用户提交的登录信息
	username := user.Username
	password := user.Password

	// 用户验证
	if username == "admin" && password == "123456" {
		// 生成JWT token
		token, err := middlewares.SetToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "生成token失败",
			})
		}
		// 登录成功
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"token":   token,
		})
	} else {
		// 登录失败
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户名或密码错误",
		})
	}
}
