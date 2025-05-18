package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	CreateProduct(c *gin.Context)
	ListProducts(c *gin.Context)
}

type handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	resp, err := h.usecase.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *handler) ListProducts(c *gin.Context) {
	products, err := h.usecase.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)

}
