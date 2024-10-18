package product_repository

import (
	"database/sql"
	"errors"
	"log"

	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pr *ProductRepository) GetById(id int64) (product_domain.ProductDomainInterface, error) {
    query := "SELECT id, name, price FROM products WHERE id = ?"
    row := pr.connection.QueryRow(query, id)

    var productId int64
    var productName string
    var productPrice float64

    if err := row.Scan(&productId, &productName, &productPrice); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            log.Printf("Produto com id %d não encontrado", id)
            return nil, nil // Retorne nil se não encontrar o produto
        }
        log.Printf("Erro ao escanear o produto: %v", err)
        return nil, err
    }

    // Cria uma instância de ProductDomain e preenche com os dados
    product := product_domain.NewProductDomain(productName, productPrice)
    product.SetID(productId)

    return product, nil
}