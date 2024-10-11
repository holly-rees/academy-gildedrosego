package repository

import (
	"gildedrose/model"
	"strings"
)

type InMemoryItemRepository struct{}

var items []*model.Item

func NewInMemoryItemRepository() *InMemoryItemRepository {
	InitDB()
	return &InMemoryItemRepository{}
}

func InitDB() {
	items = []*model.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}, // <-- :O
	}
}

func GildedItemFactory(item *model.Item) model.GildedItem {
	switch {
	case strings.Contains(item.Name, "Backstage pass"):
		return model.NewBackstagePassItem(item)

	case item.Name == "Sulfuras, Hand of Ragnaros":
		return model.NewSulfurasItem(item)

	case item.Name == "Aged Brie":
		return model.NewAgedBrieItem(item)

	case strings.Contains(item.Name, "Conjured"):
		return model.NewConjuredItem(item)

	default:
		return model.NewStandardItem(item)
	}
}

func (repo *InMemoryItemRepository) UpdateQuality() {

	for _, item := range items {

		gildedItem := GildedItemFactory(item)
		gildedItem.Update()

	}
}
