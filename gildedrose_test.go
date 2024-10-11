package main

import (
	"gildedrose/model"
	"gildedrose/repository"
	"gildedrose/service"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func Test_SulfurasQualityDoesntChange(t *testing.T) {
	var items = []*model.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	var expectedQuality = []int{80, 80}

	UpdateQuality(items)
	UpdateQuality(items)

	for i, item := range items {
		if item.Quality != expectedQuality[i] {
			t.Errorf("Item %d quality not equal. Got %d, expected %d", i, item.Quality, expectedQuality[i])
		}
	}
}

func Test_SulfurasDoesntChange(t *testing.T) {
	var items = []*model.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	var expectedItems = []*model.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	UpdateQuality(items)
	UpdateQuality(items)

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item SellIn: got %v, expected %v", items[0].SellIn, expectedItems[0].SellIn)
	}

}

func Test_BackstagePassQualityIncreaseBy1Before11DaySellIn(t *testing.T) {
	var items = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	}

	var expectedItems = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
	}

	UpdateQuality(items)
	UpdateQuality(items)

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_BackstagePassQualityIncreaseBy2After11DaySellIn(t *testing.T) {
	var items = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 12, 20},
	}

	var expectedItems = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 9, 24},
	}

	UpdateQuality(items)
	UpdateQuality(items)
	UpdateQuality(items)

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_BackstagePassQualityIncreaseBy3After6DaySellIn(t *testing.T) {
	var items = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
	}

	var expectedItems = []*model.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 4, 38},
	}

	for i := 0; i < 8; i++ {
		UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_AgedBrieQuality(t *testing.T) {
	var items = []*model.Item{
		{"Aged Brie", 2, 0},
	}

	var expectedItems = []*model.Item{
		{"Aged Brie", -2, 6},
	}

	for i := 0; i < 4; i++ {
		UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}
}

func Test_StandardItemsQuality(t *testing.T) {
	var items = []*model.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Elixir of the Mongoose", 5, 7},
	}
	var expectedItems = []*model.Item{
		{"+5 Dexterity Vest", -2, 6},
		{"Elixir of the Mongoose", -7, 0},
	}

	for i := 0; i < 12; i++ {
		UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}
}

func Test_SellInDecreases(t *testing.T) {
	var items = []*model.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	}

	var expectedSellIns = []int{
		-3,
		-11,
		-8,
		0,
		-1,
		2,
		-3,
		-8,
	}

	for i := 0; i < 13; i++ {
		UpdateQuality(items)
	}

	for i, item := range items {
		if item.SellIn != expectedSellIns[i] {
			t.Errorf("Item %d SellIns not equal. Got %d, expected %d", i, item.SellIn, expectedSellIns[i])
		}
	}
}

func Test_ConjuredItemsQuality(t *testing.T) {
	var items = []*model.Item{
		{"Conjured Mana Cake", 3, 6},
		{"Conjured Mana Cake", 5, 10},
	}

	var expectedItems = []*model.Item{
		{"Conjured Mana Cake", 0, 0},
		{"Conjured Mana Cake", 2, 4},
	}

	for i := 0; i < 3; i++ {
		UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item sellin, quality: got %v, %v, expected %v, %v", items[0].SellIn, items[0].Quality, expectedItems[0].SellIn, expectedItems[0].Quality)
	}

}

func Test_Characterisation_15_Day_Output(t *testing.T) {
	t.Skip("added new features so skipping original characterisation test")

	cmd := exec.Command("go", "run", "../texttest_fixture.go", "15")
	actualOutput, _ := cmd.Output()

	expectedOutput, err := os.ReadFile("../expected-15-days-output.txt")
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	if string(actualOutput) != string(expectedOutput) {
		t.Errorf("Output does not match.\nExpected: %s\nGot: %s", expectedOutput, actualOutput)
	}
}

func TestGetItemsFromRepository(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryItemRepository()
	service := service.NewItemService(*repo)
	expected := []*model.Item{
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

	// Act
	got, err := service.GetItems()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestUpdateItemsInRepository(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryItemRepository()
	service := service.NewItemService(*repo)
	expected := []*model.Item{
		{Name: "+5 Dexterity Vest", SellIn: 9, Quality: 19},
		{Name: "Aged Brie", SellIn: 1, Quality: 1},
		{Name: "Elixir of the Mongoose", SellIn: 4, Quality: 6},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 14, Quality: 21},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 50},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 50},
		{Name: "Conjured Mana Cake", SellIn: 2, Quality: 4}, // <-- :O
	}

	// Act
	err := service.UpdateQuality()
	if err != nil {
		t.Errorf("Unexpected error updating items: %v", err)
	}

	got, err := service.GetItems()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
