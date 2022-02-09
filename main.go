package main

import (
	"go-crud/controllers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Halo"})
	})

	r.GET("/tasks", controllers.FindTasks)
	r.GET("/tasks/raw", controllers.FindTasksRaw)
	r.POST("/task", controllers.CreateTask)
	r.GET("/task/:id", controllers.FindTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.Run(":5000")
}
