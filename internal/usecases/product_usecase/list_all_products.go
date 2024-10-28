package product_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pu *ProductUseCase) ListAll(limit, offset int) ([]product_domain.ProductDomainInterface, error) {
	products, err := pu.Repository.ListAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}