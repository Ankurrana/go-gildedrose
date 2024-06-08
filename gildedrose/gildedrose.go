package gildedrose

import "strings"

type ItemType string

var ItemDecoratorInstance ItemDecorator

func init() {
	ItemDecoratorInstance = GetItemDecorator()
}

const (
	Conjured ItemType = "conjured"
	Other    ItemType = "other"
	Passes   ItemType = "passes"
	Brie     ItemType = "brie"
	Sulfuras ItemType = "sulfuras"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemAdapter struct {
	*Item
	Type           ItemType
	QualityUpdater QualityUpdater
}

func (i *ItemAdapter) UpdateQuality() {
	i.Quality = i.QualityUpdater.GetUpdatedQuality(i.Quality, i.SellIn)
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		items[i].SellIn -= 1
		if items[i].SellIn < 0 {
			items[i].SellIn = 0
		}
		item := ItemDecoratorInstance.GetAdaptedItem(items[i])
		item.UpdateQuality()
	}
}

func (i Item) GetType() ItemType {
	name := i.Name
	if strings.Contains(name, "Conjured") {
		return Conjured
	}
	if strings.Contains(name, "Backstage passes") {
		return Passes
	}
	if strings.Contains(name, "Sulfuras") {
		return Sulfuras
	}
	if strings.Contains(name, "Aged Brie") {
		return Brie
	}
	return Other
}
