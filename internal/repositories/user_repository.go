package repositories

import (
	"database/sql"

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


// func (r *ProductRepository) Create(product *models.Product) error {
// 	_, err := r.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
// 	return err
// }

// func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
// 	var product models.Product
// 	row := r.DB.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
// 	err := row.Scan(&product.ID, &product.Name, &product.Price)
// 	return &product, err
// }

// Demais métodos para Update e Delete
