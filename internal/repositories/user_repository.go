package repositories

import (
	"database/sql"
	"errors"
	"log"
	"time"

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
	var createdAt []byte
	var updatedAt []byte

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &createdAt, &updatedAt, &user.CreatedBy, &user.UpdatedBy); err != nil {
			log.Printf("Erro ao escanear o usuário: %v", err)
			return nil, err
		}
		// Converte os valores []byte para time.Time
		var err error
		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, errors.New("erro ao converter created_at para time.Time")
		}

		user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		if err != nil {
			return nil, errors.New("erro ao converter updated_at para time.Time")
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

// GetProductById retorna um produto específico com base no ID fornecido
func (u *UserRepository) Show(id int64) (*models.User, error) {
	query := "SELECT id, username, created_at, updated_at, created_by, updated_by FROM users WHERE id = ?"
	row := u.connection.QueryRow(query, id)

	var user models.User
	var createdAt []byte
	var updatedAt []byte
	
	if err := row.Scan(&user.ID, &user.Username, &createdAt, &updatedAt, &user.CreatedBy, &user.UpdatedBy); err != nil {
		log.Printf("Erro ao escanear o usuário: %v", err)
		return nil, err
	}
	// Converte os valores []byte para time.Time
	var err error
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, errors.New("erro ao converter created_at para time.Time")
	}

	user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return nil, errors.New("erro ao converter updated_at para time.Time")
	}

	return &user, nil
}
