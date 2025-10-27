package controllers

import (
	"net/http"

	"github.com/Alitadie/fengfengzhidaodoc/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		db.Preload("Chapters.Documents").Find(&books)
		c.JSON(http.StatusOK, books)
	}
}
