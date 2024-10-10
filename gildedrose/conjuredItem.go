package gildedrose

type ConjuredItem struct {
	item *Item
}

func NewConjuredItem(item *Item) ConjuredItem {
	return ConjuredItem{item: item}
}

func (conjuredItem ConjuredItem) updateItemQuality() {
	conjuredItem.item.decreaseQualityOfItemBy(2)
}

func (conjuredItem ConjuredItem) updateItemSellIn() {
	conjuredItem.item.decreaseSellInBy1()
}

func (conjuredItem ConjuredItem) Update() {
	conjuredItem.updateItemQuality()
	conjuredItem.updateItemSellIn()
}
