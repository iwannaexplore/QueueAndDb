package commands

import (
	"QueueAndDb/pkg/kafka"
	"QueueAndDb/pkg/models"
	"fmt"
	"log"
	"time"
)

type ICommand interface {
	Execute() error
}

type generateItems struct {
	AmountOfItems int
	Kafka         kafka.IKafkaProducer
	Topic         string
}
type generateItemsWithDelay struct {
	Kafka         kafka.IKafkaProducer
	AmountOfItems int
	Delay         int
	Topic         string
}

func NewGenerateItems(amountOfItems int, topic string, kafka kafka.IKafkaProducer) ICommand {
	return &generateItems{
		AmountOfItems: amountOfItems,
		Kafka:         kafka,
		Topic:         topic,
	}
}
func NewGenerateItemsWithDelay(amountOfItems int, delay int, topic string, kafka kafka.IKafkaProducer) ICommand {
	return &generateItemsWithDelay{
		AmountOfItems: amountOfItems,
		Delay:         delay,
		Kafka:         kafka,
		Topic:         topic,
	}
}

func (g *generateItems) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		err := g.Kafka.SendMessageToPartitionInTopic(g.Topic, models.NewItem(index))
		if err != nil {
			log.Fatalf("Error sending message to partition %d: %v", index, err)
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}

func (g generateItemsWithDelay) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		time.Sleep(time.Duration(g.Delay) * time.Millisecond)
		err := g.Kafka.SendMessageToPartitionInTopic(g.Topic, models.NewItem(index))
		if err != nil {
			return err
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}
