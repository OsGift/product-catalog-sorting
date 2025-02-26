package main

import (
	"fmt"

	"github.com/OsGift/product-catalog-sorting/models"
	"github.com/OsGift/product-catalog-sorting/sorters"
)

func main() {
	products := []models.Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	context := sorters.SortContext{}

	// Sort by price
	fmt.Println("🔽 Sorted by Price:")
	context.SetStrategy(sorters.PriceSorter{})
	context.Execute(products)
	printProducts(products)

	// Sort by Sales/View ratio
	fmt.Println("\n🔽 Sorted by Sales/View Ratio:")
	context.SetStrategy(sorters.RatioSorter{})
	context.Execute(products)
	printProducts(products)
}

func printProducts(products []models.Product) {
	for _, p := range products {
		fmt.Printf("%d: %s | $%.2f | Ratio: %.4f\n", p.ID, p.Name, p.Price, p.SalesViewRatio())
	}
}
