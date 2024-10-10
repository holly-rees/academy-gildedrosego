package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

var maxQualityValue int = 50
var minQualityValue int = 0

func UpdateQuality(items []*Item) {

	for _, item := range items {

		if isStandardItem(item) {
			updateStandardItemQuality(item)
		} else {
			updateSpecialItemQuality(item)
		}

		updateSellInForItem(item)
	}

}

func updateSellInForItem(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn--
	}
}

func isStandardItem(item *Item) bool {
	return item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" && item.Name != "Sulfuras, Hand of Ragnaros"
}

func updateStandardItemQuality(item *Item) {
	switch {
	case item.SellIn <= 0:
		decreaseQualityOfItemBy(item, 2)
	default:
		decreaseQualityOfItemBy(item, 1)
	}
}

func updateSpecialItemQuality(item *Item) {
	switch {
	case item.Name == "Backstage passes to a TAFKAL80ETC concert":
		handleBackstagePassQuality(item)
	case item.Name == "Sulfuras, Hand of Ragnaros":
		item.Quality = 80
	case item.Name == "Aged Brie":
		if item.SellIn <= 0 {
			increaseQualityOfItemBy(item, 2)
		} else {
			increaseQualityOfItemBy(item, 1)
		}

	default:
		increaseQualityOfItemBy(item, 1)
	}
}

func handleBackstagePassQuality(item *Item) {
	switch {
	case item.SellIn <= 0:
		item.Quality = 0
	case item.SellIn >= 0 && item.SellIn < 6:
		increaseQualityOfItemBy(item, 3)
	case item.SellIn >= 6 && item.SellIn < 11:
		increaseQualityOfItemBy(item, 2)
	default:
		increaseQualityOfItemBy(item, 1)
	}
}

func increaseQualityOfItemBy(item *Item, increase int) {
	item.Quality = min(item.Quality+increase, maxQualityValue)
}

func decreaseQualityOfItemBy(item *Item, decrease int) {
	item.Quality = max(item.Quality-decrease, minQualityValue)
}
