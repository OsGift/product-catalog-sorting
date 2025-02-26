package sorters

import "github.com/OsGift/product-catalog-sorting/models"

/*
Sorter interface defines the contract for sorting strategies.

Why use an interface?
- It allows adding new sorting strategies without modifying existing code (Open/Closed Principle).
*/
type Sorter interface {
	Sort(products []models.Product)
}

/*
SortContext holds a reference to a sorting strategy and allows changing it dynamically.

Functions:
- SetStrategy: Assigns a new sorting strategy.
- Execute: Applies the selected sorting strategy.
*/
type SortContext struct {
	Strategy Sorter
}

/*
SetStrategy assigns a sorting strategy to the context.
This allows runtime switching of sorting behavior.
*/
func (c *SortContext) SetStrategy(s Sorter) {
	c.Strategy = s
}

/*
Execute runs the currently selected sorting strategy.
*/
func (c *SortContext) Execute(products []models.Product) {
	if c.Strategy != nil {
		c.Strategy.Sort(products)
	}
}
