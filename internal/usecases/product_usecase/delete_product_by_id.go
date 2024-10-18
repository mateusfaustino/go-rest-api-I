package product_usecase

func (pu *ProductUseCase) DeleteById(id int64) error {
	
	// Verifique se o produto existe, se necessário
	_, err := pu.Repository.GetById(id)
	if err != nil {
		return err
	}

	return pu.Repository.DeleteById(id)
}