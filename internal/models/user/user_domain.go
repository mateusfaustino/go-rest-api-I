package user_domain

import "time"

// Definição da interface
type UserDomainInterface interface {
	// Getters
	GetID() int64
	GetUsername() string
	GetPassword() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetCreatedBy() *int
	GetUpdatedBy() *int
	GetStatus() string

	// Setters
	SetID(id int64)
	SetUsername(username string)
	SetPassword(password string)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	SetCreatedBy(createdBy *int)
	SetUpdatedBy(updatedBy *int)
	SetStatus(status string)
}

// Definição da struct UserDomain
type UserDomain struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	CreatedBy *int
	UpdateBy  *int
	Status    string
}

func NewUserDomain(username string, password string) UserDomainInterface {
	return &UserDomain{
		Username: username,
		Password: password,
	}
}
func NewUserResponseDomain(id int64, username string, createdAt time.Time, updateAt time.Time, createdBy *int, updateBy *int, status string) UserDomainInterface {
	return &UserDomain{
		ID:        id,
		Username:  username,
		CreatedAt: createdAt,
		UpdateAt:  updateAt,
		CreatedBy: createdBy,
		UpdateBy:  updateBy,
		Status:    status,
	}
}

// Implementação dos getters
func (u *UserDomain) GetID() int64 {
	return u.ID
}

func (u *UserDomain) GetUsername() string {
	return u.Username
}

func (u *UserDomain) GetPassword() string {
	return u.Password
}

func (u *UserDomain) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserDomain) GetUpdatedAt() time.Time {
	return u.UpdateAt
}

func (u *UserDomain) GetCreatedBy() *int {
	return u.CreatedBy
}

func (u *UserDomain) GetUpdatedBy() *int {
	return u.UpdateBy
}

func (u *UserDomain) GetStatus() string {
	return u.Status
}

// Implementação dos setters
func (u *UserDomain) SetID(id int64) {
	u.ID = id
}

func (u *UserDomain) SetUsername(username string) {
	u.Username = username
}

func (u *UserDomain) SetPassword(password string) {
	u.Password = password
}

func (u *UserDomain) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *UserDomain) SetUpdatedAt(updatedAt time.Time) {
	u.UpdateAt = updatedAt
}

func (u *UserDomain) SetCreatedBy(createdBy *int) {
	u.CreatedBy = createdBy
}

func (u *UserDomain) SetUpdatedBy(updatedBy *int) {
	u.UpdateBy = updatedBy
}

func (u *UserDomain) SetStatus(status string) {
	u.Status = status
}
