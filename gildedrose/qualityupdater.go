package gildedrose

type QualityUpdater interface {
	GetUpdatedQuality(quantity int, sellin int) int
}

type LinearQuantityUpdater struct {
	Change int
	Max    int
}

func (l *LinearQuantityUpdater) GetUpdatedQuality(q int, sellin int) int {
	c := l.Change
	if sellin < 0 {
		c = 2 * c
	}
	q += c
	if q > l.Max {
		q = l.Max
	}
	if q < 0 {
		q = 0
	}
	return q
}

func GetLinearQuantityUpdater(c int) QualityUpdater {
	return &LinearQuantityUpdater{c, 50}
}

func GetConstantQuality() QualityUpdater {
	return &LinearQuantityUpdater{0, 1 << 31}
}

type PassesQualityUpdater struct {
	Max int
}

func (l *PassesQualityUpdater) GetUpdatedQuality(q int, sellin int) int {
	var c int = 0
	switch {
	case sellin > 10:
		c = 1
	case sellin <= 10 && sellin > 5:
		c = 2
	case sellin <= 5 && sellin > 0:
		c = 3
	case sellin <= 0:
		q = 0
		c = 0
	}
	q += c
	if q >= l.Max {
		q = l.Max
	}
	if q < 0 {
		q = 0
	}
	return q
}

func GetPassesQualityUpdater() *PassesQualityUpdater {
	return &PassesQualityUpdater{50}
}
