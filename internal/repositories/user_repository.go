package repositories

import (
	"database/sql"
	"log"

	// "github.com/99designs/gqlgen/integration/server/models-go"
	"github.com/mateusfaustino/go-rest-api-i/pkg/models"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := ur.connection.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Index(limit, offset int) ([]models.User, error) {
	query := "SELECT id, username, created_at, updated_at, created_by, updated_by FROM users LIMIT ? OFFSET ?"
	rows, err := u.connection.Query(query, limit, offset)
	if err != nil {
		log.Printf("Erro ao executar a query Index: %v", err)
		return nil, err
	}
	defer rows.Close()

	var userList []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy); err != nil {
			log.Printf("Erro ao escanear o usuário: %v", err)
			return nil, err
		}
		userList = append(userList, user)
	}

	// Verifica por erros pós-iterações, como fechamento do cursor
	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração das rows: %v", err)
		return nil, err
	}

	return userList, nil
}
