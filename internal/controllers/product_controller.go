package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		ProductUseCase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context){
	products, err:= p.ProductUseCase.GetProducts()

	if err !=nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusOK, products)
}

// func (c *ProductController) CreateProduct(ctx *gin.Context) {
// 	var product models.Product
// 	if err := ctx.ShouldBindJSON(&product); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := c.UseCase.CreateProduct(&product); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, product)
// }

// Outros métodos para Get, Update e Delete
