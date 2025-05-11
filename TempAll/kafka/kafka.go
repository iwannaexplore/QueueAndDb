package kafka

import "QueueAndDb/TempAll/models"

type IKafka interface {
	SendMessageToPartitionInTopic(topic string, partition string, item models.Item) error
}

type KafkaClient struct {
}

func NewKafkaClient() IKafka {
	return &KafkaClient{}
}

func (kc *KafkaClient) SendMessageToPartitionInTopic(topic string, partition string, item models.Item) error {
	return nil
}
