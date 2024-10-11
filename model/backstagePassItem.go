package model

type BackstagePassItem struct {
	item *Item
}

func NewBackstagePassItem(item *Item) BackstagePassItem {
	return BackstagePassItem{item: item}
}

func (backstagePass BackstagePassItem) updateItemQuality() {
	switch {
	case backstagePass.item.SellIn <= 0:
		backstagePass.item.Quality = 0
	case backstagePass.item.SellIn >= 0 && backstagePass.item.SellIn < 6:
		backstagePass.item.increaseQualityOfItemBy(3)
	case backstagePass.item.SellIn >= 6 && backstagePass.item.SellIn < 11:
		backstagePass.item.increaseQualityOfItemBy(2)
	default:
		backstagePass.item.increaseQualityOfItemBy(1)
	}
}

func (backstagePass BackstagePassItem) updateItemSellIn() {
	backstagePass.item.decreaseSellInBy1()
}

func (backstagePass BackstagePassItem) Update() {
	backstagePass.updateItemQuality()
	backstagePass.updateItemSellIn()
}
