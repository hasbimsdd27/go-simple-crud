package controllers

import (
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}

func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks, "status": "success"})
}

func FindTasksRaw(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var tasks []models.Task
	db.Raw("SELECT id, assigned_to, task FROM tasks").Scan(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks, "status": "success"})
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	task := models.Task{AssignedTo: input.AssignedTo, Task: input.Task}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindTask(c *gin.Context) {
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found", "status": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": task})
}

func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "data not found"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	db.Model(&task).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": task, "status": "success"})
}

func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Task

	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "data not found"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": true})
}
