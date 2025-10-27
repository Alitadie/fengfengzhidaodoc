package router

import (
	"time"

	"github.com/Alitadie/fengfengzhidaodoc/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gin.SetMode(gin.DebugMode) // Debug 模式会打印日志
	// gin.SetMode(gin.ReleaseMode) // Release 模式不会打印日志

	r.POST("/login", controllers.Login(db))
	r.GET("/books", controllers.GetBooks(db))
	return r
}
