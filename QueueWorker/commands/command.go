package commands

import (
	"QueueAndDb/pkg/kafka"
	"QueueAndDb/pkg/models"
	"fmt"
	"time"
)

type ICommand interface {
	Execute() error
}

type generateItems struct {
	AmountOfItems int
	Kafka         kafka.IKafkaProducer
}
type generateItemsWithDelay struct {
	Kafka         kafka.IKafkaProducer
	AmountOfItems int
	Delay         int
}

func NewGenerateItems(amountOfItems int, kafka kafka.IKafkaProducer) ICommand {
	return &generateItems{
		AmountOfItems: amountOfItems,
		Kafka:         kafka,
	}
}
func NewGenerateItemsWithDelay(amountOfItems int, delay int, kafka kafka.IKafkaProducer) ICommand {
	return &generateItemsWithDelay{
		AmountOfItems: amountOfItems,
		Delay:         delay,
		Kafka:         kafka,
	}
}

func (g *generateItems) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		err := g.Kafka.SendMessageToPartitionInTopic("", models.NewItem(index))
		if err != nil {
			return err
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}

func (g generateItemsWithDelay) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		time.Sleep(time.Duration(g.Delay) * time.Second)
		err := g.Kafka.SendMessageToPartitionInTopic("", models.NewItem(index))
		if err != nil {
			return err
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}
