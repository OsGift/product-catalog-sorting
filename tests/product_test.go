package tests

import (
	"testing"

	"github.com/OsGift/product-catalog-sorting/models"
)

func TestSalesViewRatio(t *testing.T) {
	tests := []struct {
		name     string
		product  models.Product
		expected float64
	}{
		{"Normal ratio", models.Product{SalesCount: 100, ViewsCount: 200}, 0.5},
		{"Zero views", models.Product{SalesCount: 100, ViewsCount: 0}, 0},
		{"Zero sales", models.Product{SalesCount: 0, ViewsCount: 200}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.product.SalesViewRatio(); got != tt.expected {
				t.Errorf("SalesViewRatio() = %v, want %v", got, tt.expected)
			}
		})
	}
}
