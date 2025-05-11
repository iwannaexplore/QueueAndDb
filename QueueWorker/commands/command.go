package commands

import (
	"fmt"
	"github.com/iwannaexplore/QueueAndDb/QueueWorker/kafka"
	"github.com/iwannaexplore/QueueAndDb/pkg/models"
	"time"
)

type ICommand interface {
	Execute() error
}

type GenerateItems struct {
	AmountOfItems int
	Kafka         kafka.IKafka
}
type GenerateItemsWithDelay struct {
	Kafka         kafka.IKafka
	AmountOfItems int
	Delay         int
}

func (g *GenerateItems) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		err := g.Kafka.SendMessageToPartitionInTopic("", "", models.NewItem(index))
		if err != nil {
			return err
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}

func (g GenerateItemsWithDelay) Execute() error {
	for index := 0; index < g.AmountOfItems; index++ {
		time.Sleep(time.Duration(g.Delay) * time.Second)
		err := g.Kafka.SendMessageToPartitionInTopic("", "", models.NewItem(index))
		if err != nil {
			return err
		}
		fmt.Printf("Item #%d sent successfully\n", index)
	}
	return nil
}
