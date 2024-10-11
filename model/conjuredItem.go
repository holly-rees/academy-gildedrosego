package model

type ConjuredItem struct {
	item *Item
}

func NewConjuredItem(item *Item) ConjuredItem {
	return ConjuredItem{item: item}
}

func (conjuredItem ConjuredItem) updateItemQuality() {
	switch {
	case conjuredItem.item.SellIn <= 0:
		conjuredItem.item.decreaseQualityOfItemBy(4)
	default:
		conjuredItem.item.decreaseQualityOfItemBy(2)
	}
}

func (conjuredItem ConjuredItem) updateItemSellIn() {
	conjuredItem.item.decreaseSellInBy1()
}

func (conjuredItem ConjuredItem) Update() {
	conjuredItem.updateItemQuality()
	conjuredItem.updateItemSellIn()
}
