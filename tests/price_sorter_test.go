package tests

import (
	"reflect"
	"testing"

	"github.com/OsGift/product-catalog-sorting/models"
	"github.com/OsGift/product-catalog-sorting/sorters"
)

func TestPriceSorter(t *testing.T) {
	products := []models.Product{
		{ID: 2, Price: 45.00},
		{ID: 1, Price: 10.00},
		{ID: 3, Price: 25.00},
	}

	expected := []models.Product{
		{ID: 1, Price: 10.00},
		{ID: 3, Price: 25.00},
		{ID: 2, Price: 45.00},
	}

	sorter := sorters.PriceSorter{}
	sorter.Sort(products)

	if !reflect.DeepEqual(products, expected) {
		t.Errorf("PriceSorter.Sort() = %v, want %v", products, expected)
	}
}
