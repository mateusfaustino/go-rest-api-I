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

func (pu *ProductUseCase) Index(limit, offset int) ([]models.Product, error) {
	products, err := pu.Repository.Index(limit, offset)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("nenhum produto encontrado")
	}

	return products, nil
}



// GetProductById busca um produto pelo ID usando o repositório
func (pu *ProductUseCase) Show(id int64) (*models.Product, error) {
	product, err := pu.Repository.Show(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("produto não encontrado")
	}

	return product, nil
}

// CreateProduct cria um novo produto usando o repositório
func (pu *ProductUseCase) Store(product *models.Product) error {
	return pu.Repository.Store(product)
}

func (pu *ProductUseCase) Update(product *models.Product) error {
	// Validar se o produto existe, se necessário
	_, err := pu.Repository.Show(product.Id)
	if err != nil {
		return err
	}

	return pu.Repository.Update(product)
}

func (pu *ProductUseCase) Destroy(id int64) error {
	// Verifique se o produto existe, se necessário
	_, err := pu.Repository.Show(id)
	if err != nil {
		return err
	}
	
	return pu.Repository.Destroy(id)
}
