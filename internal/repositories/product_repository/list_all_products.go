package product_repository

import (
	"log"

	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pr *ProductRepository) ListAll(limit, offset int) ([]product_domain.ProductDomainInterface, error) {
	query := "SELECT id, name, price FROM products LIMIT ? OFFSET ?"
	rows, err := pr.connection.Query(query, limit, offset)
	if err != nil {
		log.Printf("Erro ao executar a query Index: %v", err)
		return nil, err
	}
	defer rows.Close()

	var productList []product_domain.ProductDomainInterface

	for rows.Next() {
		var id int64
		var name string
		var price float64

		if err := rows.Scan(&id, &name, &price); err != nil {
			log.Printf("Erro ao escanear o produto: %v", err)
			return nil, err
		}

		// Cria um novo produto usando o construtor e seta o ID
		product := product_domain.NewProductDomain(name, price)
		product.SetID(id)

		// Adiciona o produto à lista como ProductDomainInterface
		productList = append(productList, product)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração das rows: %v", err)
		return nil, err
	}

	return productList, nil
}
