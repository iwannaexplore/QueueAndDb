package main

import (
	"QueueAndDb/QueueWorker/commands"
	"QueueAndDb/pkg/kafka"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Request struct {
	CommandName   string `json:"commandName"`
	AmountOfItems int    `json:"amountOfItems"`
	Topic         string `json:"topic"`
	Delay         int    `json:"delay"`
}

func main() {
	router := gin.Default()
	kafkaClient, err := kafka.NewProducerClient()
	if err != nil {
		log.Fatalf("Failed to create producer client: %v", err)
	}

	router.POST("/request", func(c *gin.Context) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		command, err := fromRequestToCommand(request, kafkaClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err = command.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Command executed successfully"})
	})

	router.Run(":8080")
}

func fromRequestToCommand(request Request, kafka kafka.IKafkaProducer) (commands.ICommand, error) {
	switch request.CommandName {
	case "generateItems":
		{
			return commands.NewGenerateItems(request.AmountOfItems, request.Topic, kafka), nil
		}
	case "generateItemsWithDelay":
		{
			return commands.NewGenerateItemsWithDelay(request.AmountOfItems, request.Delay, request.Topic, kafka), nil
		}
	default:
		return nil, errors.New("invalid command")
	}
}
