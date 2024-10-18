package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mateusfaustino/go-rest-api-i/internal/models"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases"
)

type UserController struct {
	UserUseCase usecases.UserUseCase
}

func NewUserController(usecase usecases.UserUseCase) UserController {
	return UserController{
		UserUseCase: usecase,
	}
}

func (p *UserController) ListAll(ctx *gin.Context) {
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
	users, err := p.UserUseCase.ListAll(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (p *UserController) GetById(ctx *gin.Context) {
	// Obtém o ID do parâmetro de rota e o converte para int64
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o UseCase para buscar o user pelo ID
	user, err := p.UserUseCase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	// Retorna o user encontrado
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
