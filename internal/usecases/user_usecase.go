package usecase

import (
	"errors"

	"github.com/mateusfaustino/go-rest-api-i/pkg/models"
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories"
)

type UserUseCase struct {
	Repository repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return UserUseCase{
		Repository: repo,
	}
}

func (pu *UserUseCase) Index(limit, offset int) ([]models.User, error) {
	users, err := pu.Repository.Index(limit, offset)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("nenhum usuário encontrado")
	}

	return users, nil
}

// GetUserById busca um produto pelo ID usando o repositório
func (uu *UserUseCase) Show(id int64) (*models.User, error) {
	user, err := uu.Repository.Show(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("produto não encontrado")
	}

	return user, nil
}