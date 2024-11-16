package routes

import (
	"github.com/Amangoyal1A/studentapi/controllers"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine, DB *gorm.DB) {
	studentGroup := router.Group("/students")
	{
		// Pass the DB instance to the controller
		studentGroup.POST("/create", controllers.CreateStudent(DB))
		studentGroup.GET("/getAll", controllers.ListStudent(DB))

	}
}
