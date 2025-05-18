package main

import (
	"fmt"
)

type Usecase interface {
	CreateProduct(input CreateProductInput) (map[string]interface{}, error)
	ListProducts() ([]Product, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateProduct(input CreateProductInput) (map[string]interface{}, error) {
	if input.OriginalPrice <= 0 {
		return nil, fmt.Errorf("Price must be > 0")
	}

	var finalPrice float64
	var applied []map[string]interface{}

	if len(input.Discounts) == 0 {
		finalPrice = input.OriginalPrice
	} else {
		finalPrice, applied = EvaluateDiscounts(input.OriginalPrice, input.Discounts)
	}

	product := &Product{
		Name:          input.Name,
		Description:   input.Description,
		OriginalPrice: input.OriginalPrice,
		FinalPrice:    finalPrice,
	}

	err := u.repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"product":          product,
		"appliedDiscounts": applied,
	}, nil
}

func (u *usecase) ListProducts() ([]Product, error) {
	products, err := u.repository.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
