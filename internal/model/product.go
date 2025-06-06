package model

import "fmt"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p Product) String() string {
	return fmt.Sprintf(
		"Product[ID: %d, Name: %s, Quantity: %d, CodeValue: %s, IsPublished: %t, Expiration: %s, Price: %.2f]",
		p.Id, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
}
