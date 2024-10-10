package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

var maxQualityValue int = 50
var minQualityValue int = 0

func (item *Item) decreaseQualityOfItemBy(decrease int) {
	item.Quality = max(item.Quality-decrease, minQualityValue)
}

func (item *Item) increaseQualityOfItemBy(increase int) {
	item.Quality = min(item.Quality+increase, maxQualityValue)
}

func (item *Item) decreaseSellInBy1() {
	item.SellIn--
}

func UpdateQuality(items []*Item) {

	for _, item := range items {

		gildedItem := GildedItemFactory(item)
		gildedItem.Update()

	}
}

func GildedItemFactory(item *Item) GildedItem {
	switch {
	case item.Name == "Backstage passes to a TAFKAL80ETC concert":
		return NewBackstagePassItem(item)

	case item.Name == "Sulfuras, Hand of Ragnaros":
		return NewSulfurasItem(item)

	case item.Name == "Aged Brie":
		return NewAgedBrieItem(item)

	case item.Name == "Conjured Mana Cake":
		return NewConjuredItem(item)

	default:
		return NewStandardItem(item)
	}
}
