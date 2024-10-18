package product_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories/product_repository"
)

type ProductUseCase struct {
	Repository product_repository.ProductRepository
}

func NewProductUseCase(repo product_repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		Repository: repo,
	}
}
