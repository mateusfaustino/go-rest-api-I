package product_repository

import (
	"log"
)

func (pr *ProductRepository) DeleteById(id int64) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := pr.connection.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar o produto: %v", err)
		return err
	}
	return nil
}