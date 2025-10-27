package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Alitadie/fengfengzhidaodoc/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("80bee19f515904a098a252c31562b86d")

// 初始化用户 如果不存在就创建 admin/123456
func InitAdminUser(db *gorm.DB) {
	var user models.User
	if err := db.First(&user, "username = ?", "admin").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
			db.Create(&models.User{Username: "admin", Password: string(hash)})
		}
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "参数错误"})
			return
		}

		var user models.User
		if err := db.First(&user, "username = ?", json.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
			return
		}
		//创建JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString(jwtKey)

		c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": tokenString})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权"})
			c.Abort()
			return
		}
		c.Next()
	}
}
