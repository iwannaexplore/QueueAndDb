package main

import (
	"QueueAndDb/DatabaseWorker/mongodb"
	"QueueAndDb/pkg/kafka"
	"fmt"
	"log"
	"time"
)

func main() {
	mongoDbClient := mongodb.NewMongoItemRepository()

	MoveFromTopicToBd(string(kafka.FirstTopic), mongoDbClient)
	MoveFromTopicToBd(string(kafka.SecondTopic), mongoDbClient)
	MoveFromTopicToBd(string(kafka.ThirdTopic), mongoDbClient)

	for {
		fmt.Println("Running...")
		time.Sleep(10 * time.Second) // Sleep to prevent high CPU usage
	}
}

func MoveFromTopicToBd(topic string, repo mongodb.ItemRepository) {
	kafkaConsumer, err := kafka.NewConsumerClient()
	//Read from kafka
	if err != nil {
		log.Fatalf("Failed to create consumer client: %v", err)
	}
	go func(topic string, kafkaClient kafka.IKafkaConsumer) {
		for {
			item, err := kafkaClient.ReadMessageFromPartitionInTopic(topic)
			if err != nil {
				log.Fatalf("Failed to read message from topic: %v", err)
			}

			log.Printf("%v\n", item)
			index, err := repo.Insert(item)
			if err != nil {
				log.Fatalf("Failed to write to db: %v", err)
			}
			log.Printf("Item #%v written to db successfully from topic %s\n", index, topic)
		}
	}(topic, kafkaConsumer)
}
