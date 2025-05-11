package models

import "fmt"

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

func NewItem(index int) Item {
	subItems := make([]SubItem, 5)
	for i := 0; i < 5; i++ {
		subItems[i] = newSubItem(index)
	}
	return Item{
		StringProperty:     fmt.Sprintf("%s-string", index),
		NumberProperty:     index,
		SubItemProperty:    newSubItem(-1),
		SubItemsProperties: subItems,
	}
}

func newSubItem(index int) SubItem {
	return SubItem{
		StringProperty: fmt.Sprintf("%s-substring", index),
		NumberProperty: index,
	}
}
