package gildedrose

type ItemDecorator struct{}

func GetItemDecorator() ItemDecorator {
	return ItemDecorator{}
}

func (f ItemDecorator) GetAdaptedItem(item *Item) ItemAdapter {
	i := ItemAdapter{Item: item}
	itemType := i.GetType()

	SimpleQDecrementor := GetLinearQuantityUpdater(-1)
	SimpleQIncrementor := GetLinearQuantityUpdater(1)
	DoubleQDecrementor := GetLinearQuantityUpdater(-2)

	i.QualityUpdater = SimpleQDecrementor

	switch itemType {
	case Conjured:
		i.QualityUpdater = DoubleQDecrementor
	case Passes:
		i.QualityUpdater = GetPassesQualityUpdater()
	case Brie:
		i.QualityUpdater = SimpleQIncrementor
	case Sulfuras:
		i.QualityUpdater = GetConstantQuality()
	}

	return i
}
