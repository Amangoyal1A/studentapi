package controllers

import (
	"net/http"
	"github.com/Amangoyal1A/studentapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateStudent(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		// Bind JSON request to struct
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save student to the database
		result := DB.Create(&student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": student})
	}
}
