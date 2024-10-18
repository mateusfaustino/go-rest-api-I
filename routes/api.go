package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api-i/internal/controllers"
	"github.com/mateusfaustino/go-rest-api-i/internal/controllers/product_controller"
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories"
	"github.com/mateusfaustino/go-rest-api-i/internal/repositories/product_repository"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases/product_usecase"
	"github.com/mateusfaustino/go-rest-api-i/internal/usecases"
)

func SetupRouter(connection *sql.DB) *gin.Engine {
	router := gin.Default()

	ProductRepository := product_repository.NewProductRepository(connection)
	productUseCase := product_usecase.NewProductUseCase(ProductRepository)
	productController := product_controller.NewProductController(productUseCase)


	UserRepository := repositories.NewUserRepository(connection)
	UserUseCase := usecases.NewUserUseCase(UserRepository)
	UserController := controllers.NewUserController(UserUseCase)

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

	productRouter := router.Group("/product")
	{
		productRouter.GET("/", productController.ListAll)
		productRouter.GET("/:id", productController.GetById)
		productRouter.POST("/", productController.Store)
		productRouter.PUT("/:id", productController.UpdateById)
		productRouter.DELETE("/:id", productController.DeleteById)
	}


	userRouter := router.Group("/user")
	{
		userRouter.GET("/", UserController.ListAll)
		userRouter.GET("/:id", UserController.GetById)
		// userRouter.POST("/", UserController.Store)
		// userRouter.PUT("/:id", UserController.UpdateById)
		// userRouter.DELETE("/:id", UserController.DeleteById)
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
