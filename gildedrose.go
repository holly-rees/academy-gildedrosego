package main

import (
	"gildedrose/model"
	"strings"
)

func UpdateQuality(items []*model.Item) {

	for _, item := range items {

		gildedItem := GildedItemFactory(item)
		gildedItem.Update()

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
