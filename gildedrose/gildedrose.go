package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

var maxQualityValue int = 50

func UpdateQuality(items []*Item) {
	for _, item := range items {

		updateSellInForItem(item)

		if isStandardItem(item) {
			updateStandardItemQuality(item)
		} else {
			updateSpecialItemQuality(item)
		}
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
	if item.SellIn < 0 {
		item.Quality = max(item.Quality-2, 0)
	} else {
		item.Quality = max(item.Quality-1, 0)
	}
}

func updateSpecialItemQuality(item *Item) {
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		handleBackstagePassQuality(item)
	} else {
		item.Quality = min(item.Quality+1, maxQualityValue)
	}
}

func handleBackstagePassQuality(item *Item) {
	switch {
	case item.SellIn < 0:
		item.Quality = 0
	case item.SellIn < 6:
		item.Quality = min(item.Quality+3, maxQualityValue)
	case item.SellIn < 11:
		item.Quality = min(item.Quality+2, maxQualityValue)
	default:
		item.Quality = min(item.Quality+1, maxQualityValue)
	}
}
