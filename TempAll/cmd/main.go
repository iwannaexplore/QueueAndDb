package main

import (
	"QueueAndDb/TempAll/commands"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	CommandName   string `json:"commandName"`
	AmountOfItems int    `json:"amountOfItems"`
	Delay         int    `json:"delay"`
}

func main() {
	router := gin.Default()

	router.POST("/request", func(c *gin.Context) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		command, err := fromRequestToCommand(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err = command.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, request)
	})

	router.Run(":8080")
}

func fromRequestToCommand(request Request) (commands.ICommand, error) {
	switch request.CommandName {
	case "GenerateItems":
		{
			return &commands.GenerateItems{AmountOfItems: request.AmountOfItems}, nil
		}
	case "GenerateItemsWithDelay":
		{
			return &commands.GenerateItemsWithDelay{AmountOfItems: request.AmountOfItems, Delay: request.Delay}, nil
		}
	default:
		return nil, errors.New("invalid command")
	}
}
