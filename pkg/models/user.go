package models

import "time"

// User representa um usuário com seus atributos correspondentes à tabela users.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=255"` // Adicionando validações
	Password  string    `json:"password" validate:"required,min=6"`         // Adicionando validações
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by,omitempty"` // Pode ser nil
	UpdatedBy *int      `json:"updated_by,omitempty"` // Pode ser nil
}
