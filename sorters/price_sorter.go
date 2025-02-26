package sorters

import (
	"sort"

	"github.com/OsGift/product-catalog-sorting/models"
)

type PriceSorter struct{}

func (p PriceSorter) Sort(products []models.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}
