package repositories

import (
	"database/sql"
	"log"
	"errors"

	"github.com/mateusfaustino/go-rest-api-i/pkg/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// Index retorna uma lista paginada de produtos com base nos parâmetros limit e offset.
func (pr *ProductRepository) Index(limit, offset int) ([]models.Product, error) {
	query := "SELECT id, name, price FROM products LIMIT ? OFFSET ?"
	rows, err := pr.connection.Query(query, limit, offset)
	if err != nil {
		log.Printf("Erro ao executar a query Index: %v", err)
		return nil, err
	}
	defer rows.Close()

	var productList []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			log.Printf("Erro ao escanear o produto: %v", err)
			return nil, err
		}
		productList = append(productList, product)
	}

	// Verifica por erros pós-iterações, como fechamento do cursor
	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração das rows: %v", err)
		return nil, err
	}

	return productList, nil
}


// GetProductById retorna um produto específico com base no ID fornecido
func (pr *ProductRepository) Show(id int64) (*models.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = ?"
	row := pr.connection.QueryRow(query, id)

	var product models.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Produto com id %d não encontrado", id)
			return nil, nil // Ou retorne um erro personalizado, se preferir
		}
		log.Printf("Erro ao escanear o produto: %v", err)
		return nil, err
	}

	return &product, nil
}

// CreateProduct insere um novo produto no banco de dados
func (pr *ProductRepository) Store(product *models.Product) error {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	result, err := pr.connection.Exec(query, product.Name, product.Price)
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
	product.Id = id

	return nil
}

// UpdateProduct atualiza os detalhes de um produto existente.
func (pr *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := pr.connection.Exec(query, product.Name, product.Price, product.Id)
	if err != nil {
		log.Printf("Erro ao atualizar o produto: %v", err)
		return err
	}
	return nil
}

// DeleteProduct remove um produto do banco de dados usando o ID.
func (pr *ProductRepository) Destroy(id int64) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar o produto: %v", err)
		return err
	}
	return nil
}
