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

func (s *ItemService) GetItems() ([]*model.Item, error) {
	items, err := s.repository.GetItems()
	if err != nil {
		return []*model.Item{}, err
	}
	return items, nil
}

func (s *ItemService) UpdateQuality() error {
	items, err := s.GetItems()
	if err != nil {
		return err
	}

	for _, item := range items {

		gildedItem := GildedItemFactory(item)
		gildedItem.Update()

	}

	return nil
}
