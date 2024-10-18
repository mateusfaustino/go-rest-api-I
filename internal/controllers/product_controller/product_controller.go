package product_controller

import "github.com/mateusfaustino/go-rest-api-i/internal/usecases/product_usecase"

type ProductController struct {
	ProductUseCase product_usecase.ProductUseCase
}

func NewProductController(usecase product_usecase.ProductUseCase) ProductController {
	return ProductController{
		ProductUseCase: usecase,
	}
}
