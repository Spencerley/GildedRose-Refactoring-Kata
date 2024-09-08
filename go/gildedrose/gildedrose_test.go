package gildedrose_test

import (
	"testing"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
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

// Test for SellIn value decreasing by one before sell in date has passed
// func Test_Sell_In_Value_Decreases_By_One(t *testing.T) {
// 	var items = []*gildedrose.Item{
// 		{"Apple", 1, 1},
// 		{"Banana", 1, 1},
// 	}

// }

// Test for quality decreasing by one
func Test_Quality_Decreases_By_One(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Apple", 1, 1},
	}
	gildedrose.UpdateQuality(items)
	println(items)

	expected := 0
	actual := items[0].Quality

	if expected != actual {
		t.Errorf("got %q, wanted %q", expected, actual)
	}
}

// Test for quality degrading by 2 after sell by date passes

// Test for quality never being negative

func Test_Quality_Not_Negative(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Apple", 3, 3},
		{"Banana", 1, 1},
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		if items[i].Quality < 0 {
			t.Errorf("Fail: quality below 0")
		} 
	}
}

// Test for Aged Brie increasing in quality

// Test for items quality never getting increased above 50
func Test_Quality_Not_Increased_Above_50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Quality > 50 {
		t.Errorf("Fail: quality above 50")
	}
}

// Test for Sulfuras never being sold

// Test for Sulfuras Quality being 80 and never changing

// Test for Backstage passes quality increases by 2 when 10 - 6 days till concert

// Test for backstage passes quality increases by 3 when 5 - 0 days to go

// Test for backstage passes quality being 0 after the concert

// Test for conjured items degrade in quality twice as fast as normal items
