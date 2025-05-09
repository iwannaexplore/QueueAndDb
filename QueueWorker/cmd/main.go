package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// GET endpoint
	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	// POST endpoint
	router.POST("/items", func(c *gin.Context) {
		var newItem Item
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		items = append(items, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	// Start the server
	router.Run(":8080")
}

//Получить запрос
//Нагенерировать данныз
//Отправить их в кафку

type IQueueRepository interface {
	GenerateItems(amountOfItems int) error
	GenerateItemsWithDelay(amountOfItems int, delay int) error
}
