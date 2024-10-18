package product_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pu *ProductUseCase) Store(product product_domain.ProductDomainInterface) error {
	return pu.Repository.Store(product)
}