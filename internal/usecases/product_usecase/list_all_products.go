package product_usecase

import (
	"errors"
	"github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pu *ProductUseCase) ListAll(limit, offset int) ([]product_domain.ProductDomainInterface, error) {
	products, err := pu.Repository.ListAll(limit, offset)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("nenhum produto encontrado")
	}

	return products, nil
}