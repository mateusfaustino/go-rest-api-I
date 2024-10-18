package product_repository

import (
	"log"

	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pr *ProductRepository) Store(product product_domain.ProductDomainInterface) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	result, err := pr.connection.Exec(query, product.GetName(), product.GetPrice())
	if err != nil {
		log.Printf("Erro ao criar produto: %v", err)
		return err
	}

	// Obtém o ID gerado e atribui ao produto
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao obter o ID do produto: %v", err)
		return err
	}
	product.SetID(id)

	return nil
}
