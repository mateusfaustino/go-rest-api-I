package user_repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	user_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/user"
)

func (pr *UserRepository) GetById(id int64) (user_domain.UserDomainInterface, error) {
	query := "SELECT id, username, created_at, updated_at, created_by, updated_by, status FROM users WHERE id = ?"
	row := pr.connection.QueryRow(query, id)

	var userId int64
	var userName string
	var createdAtStr string // Receber como string
	var updatedAtStr string // Receber como string
	var createdBy *int
	var updatedBy *int
	var status string

	if err := row.Scan(&userId, &userName, &createdAtStr, &updatedAtStr, &createdBy, &updatedBy, &status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Produto com id %d não encontrado", id)
			return nil, nil // Retorne nil se não encontrar o produto
		}
		log.Printf("Erro ao escanear o produto: %v", err)
		return nil, err
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		log.Printf("Erro ao converter created_at: %v", err)
		return nil, err
	}

	updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr)
	if err != nil {
		log.Printf("Erro ao converter updated_at: %v", err)
		return nil, err
	}

	// Cria uma instância de ProductDomain e preenche com os dados
	user := user_domain.NewUserResponseDomain(
		userId,
		userName,
		createdAt,
		updatedAt,
		createdBy,
		updatedBy,
		status,
	)

	return user, nil
}
