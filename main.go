package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var e error

type Task struct {
	ID      uint   `json:"id"`
	Title    string `json:"title"`
	Description string `json:"description"`
  Date Time `json:"date"`
  Status uint8 `json:"status"`
}

func main() {
	db, e = gorm.Open("sqlite3", "./example.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	r := gin.Default()
	// Get tasks
	r.GET("/tasks", getTasks)
	// Get task by id
	r.GET("/task/:id", getTaskById)
	// Insert new customer
	r.POST("/tasks", saveTask)
	// Update task
	r.PUT("/tasks/:id", updateTask)
	// Delete task
	r.DELETE("/tasks/:id", deleteTask)
	r.Run(":1991")
}

// Get tasks
func getTasks(c *gin.Context) {
	var tasks []Task
	if e := db.Find(&tasks).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, tasks)
	}
}

// Get task by id
func getTaskById(c *gin.Context) {
	var task Task
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&task).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, task)
	}
}

// Insert new task
func saveTask(c *gin.Context) {
	var task Task
	c.BindJSON(&task)
	db.Create(&task)
	c.JSON(200, task)
}

// Update task
func updateTask(c *gin.Context) {
	var task Task
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&task).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.BindJSON(&task)
		db.Save(&task)
		c.JSON(200, task)
	}
}

// Delete task
func deleteTask(c *gin.Context) {
	var task Task
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&task)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
