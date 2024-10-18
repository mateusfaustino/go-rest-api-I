package usecases

import (
	"errors"

	"github.com/mateusfaustino/go-rest-api-i/internal/models"
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

func (pu *UserUseCase) ListAll(limit, offset int) ([]models.User, error) {
	users, err := pu.Repository.ListAll(limit, offset)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("nenhum usuário encontrado")
	}

	return users, nil
}

// GetUserById busca um produto pelo ID usando o repositório
func (uu *UserUseCase) GetById(id int64) (*models.User, error) {
	user, err := uu.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("produto não encontrado")
	}

	return user, nil
}
