package models

// Product representa um produto com id, nome e preço.
type Product struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name" validate:"required,min=3"`
	Price float64 `json:"price" validate:"required,gt=0"`
}
