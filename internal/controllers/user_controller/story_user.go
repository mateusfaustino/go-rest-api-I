package user_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	user_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/user"
)

type storeUserRequest struct {
	ID            int64       `json:"id"`
	Username      string    `json:"username" validate:"required,min=3,max=255"` // Adicionando validações
	Password      string    `json:"password" validate:"required,min=6"`         // Adicionando validações
}

type storeUserResponse struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"` 
	CreatedAt     time.Time `json:"created_at"`
	UpdateById time.Time `json:"updated_at"`
	CreatedBy     *int      `json:"created_by`
	Status string      `json:"status"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (uc *UserController) Store(ctx *gin.Context) {
	var userRequest storeUserRequest

	// Faz o bind do JSON no modelo Product
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Executa a validação
	if err := validate.Struct(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validação dos dados falhou", "details": err.Error()})
		return
	}

	var userDomain user_domain.UserDomainInterface = user_domain.NewUserDomain(
		userRequest.Username,
		userRequest.Password,
	)

	// Chama o UseCase para criar o produto
	if err := uc.UserUseCase.Store(userDomain); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	var userResponse storeUserResponse = storeUserResponse{
		ID: userDomain.GetID(),
		Username: userDomain.GetUsername(), 
		CreatedAt: userDomain.GetCreatedAt(),
		UpdateById: userDomain.GetUpdatedAt(),
		CreatedBy: userDomain.GetCreatedBy(),
		Status: userDomain.GetStatus(),
	}

	// Retorna o produto criado
	ctx.JSON(http.StatusCreated, gin.H{"user": userResponse})
}
