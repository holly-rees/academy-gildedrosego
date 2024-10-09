package gildedrose_test

import (
	"os"
	"os/exec"
	"testing"
	// "github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

// func Test_Foo(t *testing.T) {
// 	var items = []*gildedrose.Item{
// 		{"foo", 0, 0},
// 	}

// 	gildedrose.UpdateQuality(items)

// 	if items[0].Name != "fixme" {
// 		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
// 	}
// }

func Test_Characterisation_15_Day_Output(t *testing.T) {
	cmd := exec.Command("go", "run", "../texttest_fixture.go", "15")
	actualOutput, _ := cmd.Output()
	//print(actualOutput)

	expectedOutput, err := os.ReadFile("../expected-15-days-output.txt")
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	if string(actualOutput) != string(expectedOutput) {
		t.Errorf("Output does not match.\nExpected: %s\nGot: %s", expectedOutput, actualOutput)
	}
}
