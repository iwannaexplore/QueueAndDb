package commands

type ICommand interface {
	Execute() error
}

type GenerateItems struct {
	AmountOfItems int
	Kafka         IKafka
}
type GenerateItemsWithDelay struct {
	AmountOfItems int
	Delay         int
}

func (g *GenerateItems) Execute() error {

}
