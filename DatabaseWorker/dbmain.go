package main

import (
	"QueueAndDb/DatabaseWorker/mongodb"
	"QueueAndDb/pkg/kafka"
	"fmt"
	"log"
	"time"
)

func main() {
	//Read from kafka
	kafkaConsumer, err := kafka.NewConsumerClient()
	if err != nil {
		log.Fatalf("Failed to create consumer client: %v", err)
	}

	mongoDbClient := mongodb.NewMongoItemRepository()

	if err != nil {
		log.Fatalf("Failed to create consumer client: %v", err)
	}
	MoveFromTopicToBd(string(kafka.FirstTopic), kafkaConsumer, mongoDbClient)
	MoveFromTopicToBd(string(kafka.SecondTopic), kafkaConsumer, mongoDbClient)
	MoveFromTopicToBd(string(kafka.ThirdTopic), kafkaConsumer, mongoDbClient)

	for {
		fmt.Println("Running...")
		time.Sleep(10 * time.Second) // Sleep to prevent high CPU usage
	}
}

func MoveFromTopicToBd(topic string, kafkaClient kafka.IKafkaConsumer, repo mongodb.ItemRepository) {
	go func(topic string, kafkaClient kafka.IKafkaConsumer) {
		item, err := kafkaClient.ReadMessageFromPartitionInTopic(topic)
		if err != nil {
			log.Fatalf("Failed to read message from topic: %v", err)
		}

		log.Printf("%v\n", item)
		index, err := repo.Insert(item)
		if err != nil {
			log.Fatalf("Failed to write to db: %v", err)
		}
		log.Printf("Item #%d written to db successfully\n", index)
	}(topic, kafkaClient)
}
