package sorters

import "github.com/OsGift/product-catalog-sorting/models"

type Sorter interface {
	Sort(products []models.Product)
}

type SortContext struct {
	Strategy Sorter
}

func (c *SortContext) SetStrategy(s Sorter) {
	c.Strategy = s
}

func (c *SortContext) Execute(products []models.Product) {
	if c.Strategy != nil {
		c.Strategy.Sort(products)
	}
}
