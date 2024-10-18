package product_usecase

import (
	"github.com/mateusfaustino/go-rest-api-i/internal/models/product"
)

func (pu *ProductUseCase) UpdateById(product product_domain.ProductDomainInterface) error {
	
	// Validar se o produto existe, se necessário
	_, err := pu.Repository.GetById(product.GetID())
	if err != nil {
		return err
	}

	return pu.Repository.UpdateById(product)
}