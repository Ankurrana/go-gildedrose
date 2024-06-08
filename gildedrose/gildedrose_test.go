package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 4},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 4},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 4},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	var expectedItems = []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"Aged Brie", 1, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Sulfuras, Hand of Ragnaros", -2, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 6},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 7},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Conjured Mana Cake", 2, 4}, // <-- :O
	}

	gildedrose.UpdateQuality(items)

	for i := 0; i < len(items); i++ {
		if items[i].Quality != expectedItems[i].Quality {
			t.Errorf("%d:ItemName: %s,Quality: Expected %d but got %d fixme\n", i, items[i].Name, expectedItems[i].Quality, items[i].Quality)
		}
	}
}
