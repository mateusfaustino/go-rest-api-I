package product_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	// "github.com/mateusfaustino/go-rest-api-i/internal/models"
	product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

type storeProductRequest struct {
	Name  string  `json:"name" validate:"required,min=3"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (p *ProductController) Store(ctx *gin.Context) {
	var productRequest storeProductRequest

	// Faz o bind do JSON no modelo Product
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Executa a validação
	if err := validate.Struct(productRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validação dos dados falhou", "details": err.Error()})
		return
	}

	var productDomain product_domain.ProductDomainInterface = product_domain.NewProductDomain(
		productRequest.Name,
		productRequest.Price,
	)

	// Chama o UseCase para criar o produto
	if err := p.ProductUseCase.Store(productDomain); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	// Retorna o produto criado
	ctx.JSON(http.StatusCreated, gin.H{"product": productDomain})
}
