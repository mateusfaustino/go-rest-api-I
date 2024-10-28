package user_repository

import (
	"log"

	user_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/user"
)

func (pr *UserRepository) Store(user user_domain.UserDomainInterface) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := pr.connection.Exec(query, user.GetUsername(), user.GetPassword())
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

	userResponse, err := pr.GetById(id)
	if err != nil {
		log.Printf("Erro ao criar produto: %v", err)
		return err
	}

	user.SetID(userResponse.GetID())
	user.SetStatus(userResponse.GetStatus())
	user.SetCreatedAt(userResponse.GetCreatedAt())
	user.SetUpdatedAt(userResponse.GetUpdatedAt())

	return nil
}
