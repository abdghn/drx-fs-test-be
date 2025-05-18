package main

import "gorm.io/gorm"

type Repository interface {
	CreateProduct(product *Product) error
	ListProducts() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProduct(product *Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) ListProducts() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
