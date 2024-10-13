package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api-i/internal/controllers"
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories"
	usecase "github.com/mateusfaustino/go-rest-api-i/internal/usecases"
)

func SetupRouter(connection *sql.DB) *gin.Engine {
	router := gin.Default()

	ProductRepository := repositories.NewProductRepository(connection)
	productUseCase := usecase.NewProductUseCase(ProductRepository)
	productController := controllers.NewProductController(productUseCase)

	// Define as rotas públicas
	router.POST("/login", func(ctx *gin.Context) {
		// Todo: valida as credenciais do usuário e gera o token JWT
	})

	// Rota de exemplo para testar o servidor
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

    // Grupo de rotas para produtos, com prefixo `/product`
	productRouter := router.Group("/product")
	{
		productRouter.GET("/", productController.Index)
        productRouter.GET("/:id", productController.Show)
        productRouter.POST("/", productController.Store)
        productRouter.PUT("/:id", productController.Update)
        productRouter.DELETE("/:id", productController.Destroy) 
	}

	// Define as rotas protegidas
	auth := router.Group("/")
	
    auth.Use(AuthMiddleware()) // Aplica o middleware de autenticação
	{
		// auth.GET("/products", productController.GetProducts)
		// auth.GET("/products/:id", productController.GetProductByID)
		// Outros endpoints protegidos
	}

	return router
}
