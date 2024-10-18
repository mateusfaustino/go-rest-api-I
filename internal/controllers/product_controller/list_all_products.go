package product_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type listAllProductResponse struct {
	Id  int64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *ProductController) ListAll(ctx *gin.Context) {
	var productsResponseList []listAllProductResponse

	// Captura os parâmetros de paginação
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10")) // Limite padrão de 10 produtos por página
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

	// Chama a função de listagem de produtos com paginação
	products, err := p.ProductUseCase.ListAll(limit, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	for _, p := range products {
		product := listAllProductResponse{
			Id: p.GetID(),
			Name: p.GetName(),
			Price: p.GetPrice(),
		}

		productsResponseList = append(productsResponseList, product)

	}

	ctx.JSON(http.StatusOK, productsResponseList)
}