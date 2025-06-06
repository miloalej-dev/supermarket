package repository

import (
	"encoding/json"
	"github.com/miloalej-dev/supermarket/internal/model"
	"os"
)

var Products []model.Product

func init() {
	data, err := os.ReadFile("products.json")

	if err != nil {
		panic("Failed to read products.json: " + err.Error())
	}

	err = json.Unmarshal(data, &Products)

	if err != nil {
		panic("Failed to unmarshal products.json: " + err.Error())
	}
}

func FindProductByID(id int) (*model.Product, error) {
	for _, product := range Products {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, nil // Return nil if not found
}

func FindProductsByPriceGreaterThan(price float64) []model.Product {
	var result []model.Product
	for _, product := range Products {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}

func AddProduct(product model.Product) {
	Products = append(Products, product)
}
