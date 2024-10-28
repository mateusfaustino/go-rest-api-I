package user_controller


import "github.com/mateusfaustino/go-rest-api-i/internal/usecases/user_usecase"

type UserController struct {
	UserUseCase user_usecase.UserUsecase
}

func NewUserController(usecase user_usecase.UserUsecase) UserController {
	return UserController{
		UserUseCase: usecase,
	}
}
