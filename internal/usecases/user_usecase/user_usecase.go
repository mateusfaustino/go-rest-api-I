package user_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories/user_repository"
)

type UserUsecase struct {
	Repository user_repository.UserRepository
}

func NewUserUseCase(repo user_repository.UserRepository) UserUsecase {
	return UserUsecase{
		Repository: repo,
	}
}
