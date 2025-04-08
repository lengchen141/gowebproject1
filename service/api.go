package service

import (
	"github.com/gin-gonic/gin"
	"gowebproject1/common"
	"gowebproject1/models"
	"net/http"
)

func GetApi(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello",
	})
}

// GetUser 获取用户信息
func GetUser(c *gin.Context) {

	var users []models.UserInfo

	common.DB.Find(&users)
	// 手动格式化日期字段
	for i := range users {
		users[i].CreatedAtStr = users[i].CreatedAt.Format("2006-01-02 15:04:05")
		users[i].UpdatedAtStr = users[i].UpdatedAt.Format("2006-01-02 15:04:05")
	}
	// 这里实现获取用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser 创建新用户
func CreateUser(c *gin.Context) {
	// 这里实现创建新用户的逻辑
	c.JSON(http.StatusCreated, gin.H{"message": "创建用户"})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	// 这里实现更新用户信息的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "更新用户信息"})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// 这里实现删除用户的逻辑
	c.JSON(http.StatusOK, gin.H{"message": "删除用户"})
}
