package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()// Connect to the database
	if err != nil {
		panic(err)
	}
 
	// Camada Repository
	ProductRepository := repository.NewProductReposity(dbConnection)

	// Camada UseCase
	ProductUseCase:= usecase.NewProductUseCase(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.Run(":8000")

}
