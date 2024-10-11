package api

import (
	"encoding/json"
	"gildedrose/model"
	"gildedrose/service"
	"net/http"
)

type APIItem struct {
	Name    string `json:"name"`
	SellIn  int    `json:"sellin"`
	Quality int    `json:"quality"`
}

type ItemAPI struct {
	itemService *service.ItemService
}

func NewItemAPI(itemService *service.ItemService) *ItemAPI {
	return &ItemAPI{
		itemService: itemService,
	}
}

func (api *ItemAPI) GetItems(writer http.ResponseWriter, request *http.Request) {
	items, err := api.itemService.GetItems()
	apiItems := ConvertToAPIItems(items)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(apiItems)
}

func (api *ItemAPI) UpdateItems(writer http.ResponseWriter, request *http.Request) {
	err := api.itemService.UpdateQuality()
	if err != nil {
		http.Error(writer, "Internal Server Error updating quality", http.StatusInternalServerError)
		return
	}
	items, err := api.itemService.GetItems()
	apiItems := ConvertToAPIItems(items)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(apiItems)
}

func ConvertToAPIItems(items []*model.Item) []APIItem {
	apiItems := make([]APIItem, len(items))
	for i, item := range items {
		apiItems[i] = APIItem{Name: item.Name,
			Quality: item.Quality,
			SellIn:  item.SellIn,
		}
	}
	return apiItems
}
