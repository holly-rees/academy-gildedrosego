package gildedrose_test

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_SulfurasQualityDoesntChange(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
	}

	var expectedQuality = []int{80, 80}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Quality != expectedQuality[i] {
			t.Errorf("Item %d quality not equal. Got %d, expected %d", i, item.Quality, expectedQuality[i])
		}
	}
}

func Test_BackstagePassQualityIncreaseBy1Before11DaySellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	}

	var expectedItems = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
	}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_BackstagePassQualityIncreaseBy2After11DaySellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 12, 20},
	}

	var expectedItems = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 9, 24},
	}

	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)
	gildedrose.UpdateQuality(items)

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_BackstagePassQualityIncreaseBy3After6DaySellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
	}

	var expectedItems = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 4, 38},
	}

	for i := 0; i < 8; i++ {
		gildedrose.UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}

}

func Test_AgedBrieQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 2, 0},
	}

	var expectedItems = []*gildedrose.Item{
		{"Aged Brie", -2, 6},
	}

	for i := 0; i < 4; i++ {
		gildedrose.UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}
}

func Test_StandardItemsQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Elixir of the Mongoose", 5, 7},
	}
	var expectedItems = []*gildedrose.Item{
		{"+5 Dexterity Vest", -2, 6},
		{"Elixir of the Mongoose", -7, 0},
	}

	for i := 0; i < 12; i++ {
		gildedrose.UpdateQuality(items)
	}

	if !reflect.DeepEqual(items, expectedItems) {
		t.Errorf("Items not equal. For first item quality: got %v, expected %v", items[0].Quality, expectedItems[0].Quality)
	}
}

func Test_Characterisation_15_Day_Output(t *testing.T) {
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
