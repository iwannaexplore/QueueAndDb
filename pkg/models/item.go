package models

type Item struct {
	StringProperty     string
	NumberProperty     int
	SubItemProperty    SubItem
	SubItemsProperties []SubItem
}

type SubItem struct {
	StringProperty string
	NumberProperty int
}

func NewItem() Item {
	return Item{
		StringProperty:     "string",
		NumberProperty:     0,
		SubItemProperty:    SubItem{},
		SubItemsProperties: []SubItem{},
	}
}
