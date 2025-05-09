package kafka

type IKafka interface {
	SendMessageToPartition(topic string, item models.Item) error
}
git remote add origin git@github.com:iwannaexplore/QueueAndDb.git
git branch -M master
git push -u origin master