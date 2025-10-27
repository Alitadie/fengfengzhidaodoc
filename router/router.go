package router

import (
	"github.com/Alitadie/fengfengzhidaodoc/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/login", controllers.Login(db))
	r.GET("/books", controllers.GetBooks(db))
	return r
}
