package main

type Product struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	OriginalPrice float64 `json:"originalPrice"`
	FinalPrice    float64 `json:"finalPrice"`
}

type DiscountInput struct {
	Type        string  `json:"type"`
	Value       float64 `json:"value,omitempty"`
	Condition   float64 `json:"condition,omitempty"`
	Tiers       []Tier  `json:"tiers,omitempty"`
	MaxDiscount float64 `json:"maxDiscount,omitempty"`
}

type Tier struct {
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Value float64 `json:"value"`
}

type CreateProductInput struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	OriginalPrice float64         `json:"originalPrice"`
	Discounts     []DiscountInput `json:"discounts"`
}
