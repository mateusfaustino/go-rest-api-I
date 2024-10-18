package product_usecase

import (
	"errors"
	"github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pu *ProductUseCase) GetById(id int64) (product_domain.ProductDomainInterface, error) {
	product, err := pu.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("produto não encontrado")
	}

	return product, nil
}