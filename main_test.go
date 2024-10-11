package main

import (
	"encoding/json"
	"gildedrose/api"
	"gildedrose/model"
	"gildedrose/repository"
	"gildedrose/service"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)

	wantCode := http.StatusOK
	wantBody := "Hello, world!"

	// Act
	handler.ServeHTTP(rr, req)

	gotCode := rr.Code
	gotBody := rr.Body.String()

	// Assert
	if gotCode != wantCode {
		t.Errorf("handler returned unexpected status code: got %v want %v", gotCode, wantCode)
	}

	if gotBody != wantBody {
		t.Errorf("handler returned unexpected body: got %v want %v", gotBody, wantBody)
	}
}

func TestGetItemsHandlerWithServer(t *testing.T) {
	// Arrange
	itemRepo := repository.NewInMemoryItemRepository()
	itemService := service.NewItemService(*itemRepo)
	itemAPI := api.NewItemAPI(itemService)
	server := httptest.NewServer(http.HandlerFunc(itemAPI.GetItems))
	defer server.Close()

	want := []model.Item{
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

	wantCode := http.StatusOK

	// Act
	resp, err := http.Get(server.URL + "/api/items")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	gotCode := resp.StatusCode
	defer resp.Body.Close()

	// Assert
	if gotCode != wantCode {
		t.Errorf("handler returned unexpected status code: got %v want %v", gotCode, wantCode)
	}

	// Check the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var got []model.Item
	if err := json.Unmarshal(bodyBytes, &got); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}
