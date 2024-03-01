package controllers

import (
	"net/http"

	"github.com/chyuhung/my-dashboard/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	// 获取用户提交的登录信息
	username := user.Username
	password := user.Password

	// 简单的验证逻辑
	if username == "admin" && password == "123456" {
		// 登录成功
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	} else {
		// 登录失败
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户名或密码错误",
		})
	}
}

func Logout(c *gin.Context) {
	// 清除用户的登录状态或相关信息
	// 这里可以根据你的实际情况进行处理，例如清除 Session、删除 Token 等

	// 登出成功
	c.JSON(http.StatusOK, gin.H{
		"message": "登出成功",
	})
}

func Register(c *gin.Context) {
	// 获取用户提交的注册信息
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 简单的验证逻辑
	if username == "" || password == "" {
		// 注册信息不完整
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "注册信息不完整",
		})
		return
	}

	// TODO: 执行注册操作，例如将用户信息保存到数据库

	// 注册成功
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}
