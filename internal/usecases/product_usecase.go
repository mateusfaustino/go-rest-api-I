package usecase

import (
	"errors"

	"github.com/mateusfaustino/go-rest-api-i/pkg/models"
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories"
)

type ProductUseCase struct {
	Repository repositories.ProductRepository
}

func NewProductUseCase(repo repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		Repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	products, err := pu.Repository.GetProducts()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("nenhum produto encontrado")
	}

	return products, nil
}


// func (uc *ProductUseCase) CreateProduct(product *models.Product) error {
// 	return uc.Repo.Create(product)
// }

// func (uc *ProductUseCase) GetProductByID(id int) (*models.Product, error) {
// 	return uc.Repo.GetByID(id)
// }

// Outros métodos para Update e Delete
