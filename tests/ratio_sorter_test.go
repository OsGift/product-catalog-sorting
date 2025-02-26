package tests

import (
	"reflect"
	"testing"

	"github.com/OsGift/product-catalog-sorting/models"
	"github.com/OsGift/product-catalog-sorting/sorters"
)

func TestRatioSorter(t *testing.T) {
	products := []models.Product{
		{ID: 1, SalesCount: 10, ViewsCount: 20},   // Ratio = 0.5
		{ID: 2, SalesCount: 100, ViewsCount: 100}, // Ratio = 1.0
		{ID: 3, SalesCount: 5, ViewsCount: 10},    // Ratio = 0.5
	}

	expected := []models.Product{
		{ID: 2, SalesCount: 100, ViewsCount: 100}, // Highest ratio
		{ID: 1, SalesCount: 10, ViewsCount: 20},
		{ID: 3, SalesCount: 5, ViewsCount: 10},
	}

	sorter := sorters.RatioSorter{}
	sorter.Sort(products)

	if !reflect.DeepEqual(products, expected) {
		t.Errorf("RatioSorter.Sort() = %v, want %v", products, expected)
	}
}
