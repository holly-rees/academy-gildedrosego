package model

type SulfurasItem struct {
	item *Item
}

func NewSulfurasItem(item *Item) SulfurasItem {
	return SulfurasItem{item: item}
}

func (sulfuras SulfurasItem) updateItemQuality() {
	sulfuras.item.Quality = 80
}

func (sulfuras SulfurasItem) Update() {
	sulfuras.updateItemQuality()
}
