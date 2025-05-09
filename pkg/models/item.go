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
