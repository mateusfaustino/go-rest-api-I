package product_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	// product_domain "github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (p *ProductController) DeleteById(ctx *gin.Context) {
	// Extrair o ID do produto da URL
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Deletar o produto
	if err := p.ProductUseCase.DeleteById(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o produto"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}