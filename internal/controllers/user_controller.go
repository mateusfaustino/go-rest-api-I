package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases"
	_"github.com/mateusfaustino/go-rest-api-i/pkg/models"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		UserUseCase: usecase,
	}
}

func (p *UserController) Index(ctx *gin.Context) {
	// Captura os parâmetros de paginação
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10")) // Limite padrão de 10 usuários por página
	if err != nil || limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'limit' inválido"})
		return
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1")) // Página padrão é a primeira
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'page' inválido"})
		return
	}

	offset := (page - 1) * limit // Calcula o deslocamento com base na página atual e no limite

	// Chama a função de listagem de usuários com paginação
	users, err := p.UserUseCase.Index(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
