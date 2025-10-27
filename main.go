package main

import (
	"log"
	"time"

	"github.com/Alitadie/fengfengzhidaodoc/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("doc.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	/*
		db.AutoMigrate(&models.User{}, &models.Book{}, &models.Chapter{}, &models.Document{})

		db.FirstOrCreate(&models.User{}, models.User{
			Username: "admin",
			Password: "123456",
		})
	*/

	controllers.InitAdminUser(db)

	r := gin.Default()

	// CORS 支持前端
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", controllers.Login(db))

	auth := r.Group("/")
	auth.Use(controllers.AuthMiddleware())
	{
		auth.GET("/books", func(c *gin.Context) { c.JSON(200, "books接口已保护") })
	}

	r.Run(":8000")
}
