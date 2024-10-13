package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases"
	"github.com/mateusfaustino/go-rest-api-i/pkg/models"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		ProductUseCase: usecase,
	}
}

func (p *ProductController) Index(ctx *gin.Context) {
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
	products, err := p.ProductUseCase.Index(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}



func (p *ProductController) Show(ctx *gin.Context) {
	// Obtém o ID do parâmetro de rota e o converte para int64
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o UseCase para buscar o produto pelo ID
	product, err := p.ProductUseCase.Show(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	// Retorna o produto encontrado
	ctx.JSON(http.StatusOK, gin.H{"product": product})
}

// CreateProduct é o handler para criar um novo produto
func (p *ProductController) Store(ctx *gin.Context) {
	var product models.Product

	// Faz o bind do JSON no modelo Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o UseCase para criar o produto
	if err := p.ProductUseCase.Store(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	// Retorna o produto criado
	ctx.JSON(http.StatusCreated, gin.H{"product": product})
}

func (p *ProductController) Update(ctx *gin.Context) {
	var product models.Product

	// Extrair o ID do produto da URL
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	product.Id = id

	// Bind do JSON do corpo da requisição para o objeto produto
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualizar o produto
	if err := p.ProductUseCase.Update(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o produto"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso"})
}


func (p *ProductController) Destroy(ctx *gin.Context) {
	// Extrair o ID do produto da URL
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Deletar o produto
	if err := p.ProductUseCase.Destroy(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o produto"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}
