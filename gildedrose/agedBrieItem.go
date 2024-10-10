package gildedrose

type AgedBrieItem struct {
	item *Item
}

func NewAgedBrieItem(item *Item) AgedBrieItem {
	return AgedBrieItem{item: item}
}

func (agedBrie AgedBrieItem) updateItemQuality() {
	if agedBrie.item.SellIn <= 0 {
		agedBrie.item.increaseQualityOfItemBy(2)
	} else {
		agedBrie.item.increaseQualityOfItemBy(1)
	}
}

func (agedBrie AgedBrieItem) updateItemSellIn() {
	agedBrie.item.decreaseSellInBy1()
}

func (agedBrie AgedBrieItem) Update() {
	agedBrie.updateItemQuality()
	agedBrie.updateItemSellIn()
}
