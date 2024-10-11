package model

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
