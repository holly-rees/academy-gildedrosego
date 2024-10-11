package service

import (
	"gildedrose/model"
	"gildedrose/repository"
	"strings"
)

type ItemService struct {
	repository repository.InMemoryItemRepository
}

func NewItemService(repo repository.InMemoryItemRepository) *ItemService {
	return &ItemService{
		repository: repo,
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

func (s *ItemService) GetItems() []*model.Item {
	return s.repository.GetItems()
}

func (s *ItemService) UpdateQuality() {

	for _, item := range s.repository.GetItems() {

		gildedItem := GildedItemFactory(item)
		gildedItem.Update()

	}
}
