package controllers

import (
	"net/http"

	"github.com/Alitadie/fengfengzhidaodoc/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Username string
			Password string
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "参数错误"})
			return
		}
		var user models.User
		if err := db.Where("username = ? AND password = ?", json.Username, json.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	}
}
