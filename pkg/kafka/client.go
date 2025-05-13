package kafka

import (
	"QueueAndDb/pkg/models"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type Topic string

const (
	broker  = "localhost:9092"
	groupID = "my_group"

	FirstTopic  Topic = "first1"
	SecondTopic Topic = "second"
	ThirdTopic  Topic = "third"
)

type IKafkaProducer interface {
	SendMessageToPartitionInTopic(topic string, item models.Item) error
}

type IKafkaConsumer interface {
	ReadMessageFromPartitionInTopic(topic string) (models.Item, error)
}
type ProducerClient struct {
	producer *kafka.Producer
}
type ConsumerClient struct {
	consumer *kafka.Consumer
}

func NewProducerClient() (IKafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}

	return &ProducerClient{
		producer: producer,
	}, nil
}

func NewConsumerClient() (IKafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &ConsumerClient{
		consumer: consumer,
	}, nil
}

// SendMessageToPartitionInTopic sends a message to a specific partition of a topic
func (kc *ProducerClient) SendMessageToPartitionInTopic(topic string, item models.Item) error {
	// Marshal the item to JSON
	value, err := json.Marshal(item)
	if err != nil {
		return err
	}

	// Create a message
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}

	// Produce the message
	err = kc.producer.Produce(message, nil)
	if err != nil {
		return err
	}

	// Wait for message delivery
	e := <-kc.producer.Events()
	if ev := e.(*kafka.Message); ev.TopicPartition.Error != nil {
		return ev.TopicPartition.Error
	}

	return nil
}

// ReadMessageFromPartitionInTopic reads a message from a specific partition of a topic
func (kc *ConsumerClient) ReadMessageFromPartitionInTopic(topic string) (models.Item, error) {
	// Assign the topic and partition to the consumer
	err := kc.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, err)
	}
	// Poll for messages
	msg, err := kc.consumer.ReadMessage(-1)
	if err != nil {
		return models.Item{}, err
	}

	// Unmarshal the message value into an Item
	var item models.Item
	if err := json.Unmarshal(msg.Value, &item); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (kc *ProducerClient) Close() {
	kc.producer.Close()
}

func (kc *ConsumerClient) Close() {
	err := kc.consumer.Close()
	if err != nil {
		log.Fatalf("Error closing consumer: %v", err)
	}
}
