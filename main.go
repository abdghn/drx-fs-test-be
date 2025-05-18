package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB()
	r := gin.Default()

	productRepository := NewRepository(db)
	productUsecase := NewUsecase(productRepository)
	productHandler := NewHandler(productUsecase)

	// POST /products to create a product
	r.POST("/products", productHandler.CreateProduct)

	// GET /products to list products
	r.GET("/products", productHandler.ListProducts)

	// Run server
	r.Run(":8080") // listen on port 8080
}
