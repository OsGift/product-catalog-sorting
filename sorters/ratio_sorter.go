package sorters

import (
	"sort"

	"github.com/OsGift/product-catalog-sorting/models"
)

type RatioSorter struct{}

func (r RatioSorter) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].SalesViewRatio() > products[j].SalesViewRatio()
	})
}
