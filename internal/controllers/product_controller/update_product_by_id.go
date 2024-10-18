package product_controller

import (
	"net/http"
	"strconv"

	// "github.com/mateusfaustino/go-rest-api-i/internal/models"
	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
	"github.com/gin-gonic/gin"
)
type updateProductRequest struct {
	Name  string  `json:"name" validate:"required,min=3"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

func (p *ProductController) UpdateById(ctx *gin.Context) {
	// var product models.Product
	var productRequest updateProductRequest

	// Extrair o ID do produto da URL
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	// product.Id = id

	// Bind do JSON do corpo da requisição para o objeto produto
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productDomain product_domain.ProductDomainInterface = product_domain.NewProductDomain(
		productRequest.Name,
		productRequest.Price,
	)
	productDomain.SetID(id)

	// Atualizar o produto
	if err := p.ProductUseCase.UpdateById(productDomain); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o produto"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso"})
}