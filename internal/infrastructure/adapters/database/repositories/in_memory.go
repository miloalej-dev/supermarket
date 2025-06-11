package repositories

import (
	"encoding/json"
	"errors"
	"github.com/miloalej-dev/supermarket/internal/application/ports/outbound"
	"github.com/miloalej-dev/supermarket/internal/domain"
	"os"
)

type InMemoryRepository struct {
	products []domain.Product
}

func NewInMemoryRepository() outbound.ProductRepository {
	m := &InMemoryRepository{}

	data, err := os.ReadFile("products.json")

	if err != nil {
		panic("failed to read products.json: " + err.Error())
	}

	err = json.Unmarshal(data, &m.products)

	if err != nil {
		panic("failed to unmarshal products.json: " + err.Error())
	}

	return m
}

// Save adds a new product to the in-memory repository.
// It appends the product to the products slice and assigns an ID based on the current length of the slice.
// It returns the saved product with its ID.
// If there is an error during the process, it returns an error.
func (i *InMemoryRepository) Save(product domain.Product) (domain.Product, error) {
	product.Id = len(i.products) + 1
	i.products = append(i.products, product)
	return product, nil
}

// FindAll retrieves all products from the in-memory repository.
// It returns a slice of Product objects.
func (i *InMemoryRepository) FindAll() []domain.Product {
	return i.products
}

// FindById retrieves a product by its ID from the in-memory repository.
// It returns a pointer to the Product object if found, or nil if not found.
func (i *InMemoryRepository) FindById(id int) *domain.Product {
	for _, product := range i.products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}

func (i *InMemoryRepository) Update(id int, product domain.Product) *domain.Product {
	for index, p := range i.products {
		if p.Id == id {
			// Update the product at the found index
			product.Id = id // Ensure the ID remains the same
			i.products[index] = product
			return &i.products[index]
		}
	}
	return nil
}

func (i *InMemoryRepository) UpdatePartial(id int, updates map[string]interface{}) *domain.Product {

	for index, p := range i.products {
		if p.Id == id {
			product := &i.products[index]
			// Update only the fields that are present in the request
			if val, ok := updates["name"]; ok {
				product.Name = val.(string)
			}
			if val, ok := updates["quantity"]; ok {
				product.Quantity = int(val.(float64))
			}
			if val, ok := updates["code_value"]; ok {
				product.CodeValue = val.(string)
			}
			if val, ok := updates["is_published"]; ok {
				product.IsPublished = val.(bool)
			}
			if val, ok := updates["expiration"]; ok {
				product.Expiration = val.(string)
			}
			if val, ok := updates["price"]; ok {
				product.Price = val.(float64)
			}

			return product
		}

	}
	return nil
}

func (i *InMemoryRepository) Delete(product domain.Product) error {
	for index, p := range i.products {
		if p.Id == product.Id {
			i.products = append(i.products[:index], i.products[index+1:]...)
			return nil
		}
	}
	return errors.New("product does not exist")
}

func (i *InMemoryRepository) FindByPriceGreaterThan(price float64) []domain.Product {
	var result []domain.Product
	for _, product := range i.products {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}
