package outbound

import "github.com/miloalej-dev/supermarket/internal/domain"

// ProductRepository defines the interface for product-related data access operations.
type ProductRepository interface {
	Repository[int, domain.Product]

	// FindByPriceGreaterThan retrieves all products with a price greater than the specified value.
	// It returns a slice of domain.Product.
	FindByPriceGreaterThan(price float64) []domain.Product
}
