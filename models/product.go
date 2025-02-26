package models

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}

func (p Product) SalesViewRatio() float64 {
	if p.ViewsCount == 0 {
		return 0
	}
	return float64(p.SalesCount) / float64(p.ViewsCount)
}
