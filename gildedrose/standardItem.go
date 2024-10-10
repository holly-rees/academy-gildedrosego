package gildedrose

type StandardItem struct {
	item *Item
}

func NewStandardItem(item *Item) StandardItem {
	return StandardItem{item: item}
}

func (standardItem StandardItem) updateItemQuality() {
	switch {
	case standardItem.item.SellIn <= 0:
		standardItem.item.decreaseQualityOfItemBy(2)
	default:
		standardItem.item.decreaseQualityOfItemBy(1)
	}
}

func (standardItem StandardItem) updateItemSellIn() {
	standardItem.item.decreaseSellInBy1()
}

func (standardItem StandardItem) Update() {
	standardItem.updateItemQuality()
	standardItem.updateItemSellIn()
}
