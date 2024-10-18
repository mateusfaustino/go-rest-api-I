package product_repository

import (
	"log"

	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pr *ProductRepository) UpdateById(product product_domain.ProductDomainInterface) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := pr.connection.Exec(query, product.GetName(), product.GetPrice(), product.GetID())
	if err != nil {
		log.Printf("Erro ao atualizar o produto: %v", err)
		return err
	}
	return nil
}
