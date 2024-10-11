package api

import (
	"encoding/json"
	"gildedrose/service"
	"net/http"
)

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
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(items)
}
