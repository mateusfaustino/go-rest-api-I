package product_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIDResponse struct {
	Id  int64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *ProductController) GetById(ctx *gin.Context) {
	
	// Obtém o ID do parâmetro de rota e o converte para int64
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o UseCase para buscar o produto pelo ID
	product, err := p.ProductUseCase.GetById(id)


	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	var productResponse GetProductByIDResponse = GetProductByIDResponse{
		Id: product.GetID(),
		Name: product.GetName(),
		Price: product.GetPrice(),
	}

	// Retorna o produto encontrado
	ctx.JSON(http.StatusOK, gin.H{"product": productResponse})
}