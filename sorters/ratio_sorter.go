package sorters

import (
	"sort"

	"github.com/OsGift/product-catalog-sorting/models"
)

/*
RatioSorter sorts products by sales/view ratio in descending order.

Why descending?
- Products with a higher sales/view ratio indicate better engagement and conversion.
*/
type RatioSorter struct{}

/*
Sort arranges products so those with higher sales/view ratios come first.
*/
func (r RatioSorter) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].SalesViewRatio() > products[j].SalesViewRatio()
	})
}
