package sorters

import (
	"sort"

	"github.com/OsGift/product-catalog-sorting/models"
)

/*
PriceSorter sorts products in ascending order of price.

Why ascending?
- Cheaper products typically attract more initial user attention.
*/
type PriceSorter struct{}

/*
Sort uses Go's built-in sort.Slice function, which sorts in O(n log n) time.

The comparator returns true if product i's price is lower than product j's.
*/
func (p PriceSorter) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}
