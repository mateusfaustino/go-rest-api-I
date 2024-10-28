package user_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/models/user"
)

func (uu *UserUsecase) Store(user user_domain.UserDomainInterface) error {
	return uu.Repository.Store(user)
}