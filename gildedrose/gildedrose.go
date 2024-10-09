package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		item := items[i]
		if isStandardItem(item) {
			decreaseQualityByOne(item)
		} else {
			if item.Quality < 50 {
				increaseQualityByOne(item)
				if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							increaseQualityByOne(item)
						}
					}
				}
			}
		}

		updateSellIn(item)

		if item.SellIn < 0 {
			if isStandardItem(item) {
				if item.Quality > 0 {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						decreaseQualityByOne(item)
					}
				}
			} else {
				increaseQualityByOne(item)
			}
		}
	}

}

func increaseQualityByOne(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func decreaseQualityByOne(item *Item) {
	if item.Quality > 0 {
		item.Quality--
	}
}

func updateSellIn(item *Item) {
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn--
	}
}

func isStandardItem(item *Item) bool {
	return item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" && item.Name != "Sulfuras, Hand of Ragnaros"
}
