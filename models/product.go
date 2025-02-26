
package models

/*
Product represents a single product in the catalog.

Fields:
- ID: Unique identifier for each product.
- Name: The name/title of the product.
- Price: Product price in USD.
- Created: Date the product was added to the catalog.
- SalesCount: Total sales of the product.
- ViewsCount: Total number of views the product page received.
*/
type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}

/*
SalesViewRatio computes how many sales occur per view.

Why this matters:
- A higher ratio means a product converts views into sales effectively.
- Handles zero views to avoid division by zero errors.
*/
func (p Product) SalesViewRatio() float64 {
	if p.ViewsCount == 0 {
		return 0
	}
	return float64(p.SalesCount) / float64(p.ViewsCount)
}
