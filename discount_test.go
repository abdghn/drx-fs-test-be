package main

import (
	"testing"
)

func TestEvaluateDiscounts(t *testing.T) {
	tests := []struct {
		name          string
		originalPrice float64
		discounts     []DiscountInput
		wantPrice     float64
		wantApplied   int
	}{
		{
			name:          "Fixed discount",
			originalPrice: 100,
			discounts: []DiscountInput{
				{Type: "fixed", Value: 20},
			},
			wantPrice:   80,
			wantApplied: 1,
		},
		{
			name:          "Percentage discount",
			originalPrice: 200,
			discounts: []DiscountInput{
				{Type: "percentage", Value: 10},
			},
			wantPrice:   180,
			wantApplied: 1,
		},
		{
			name:          "Conditional discount met",
			originalPrice: 250,
			discounts: []DiscountInput{
				{Type: "conditional", Condition: 200, Value: 30},
			},
			wantPrice:   220,
			wantApplied: 1,
		},
		{
			name:          "Tiered discount applied",
			originalPrice: 300,
			discounts: []DiscountInput{
				{Type: "tiered", Tiers: []Tier{
					{Min: 0, Max: 100, Value: 5},
					{Min: 101, Max: 299, Value: 15},
					{Min: 300, Max: 9999, Value: 25},
				}},
			},
			wantPrice:   275,
			wantApplied: 1,
		},
		{
			name:          "Cap discount applies",
			originalPrice: 300,
			discounts: []DiscountInput{
				{Type: "fixed", Value: 50},
				{Type: "percentage", Value: 20}, // 250 * 0.20 = 50
				{Type: "cap", MaxDiscount: 80},
			},
			wantPrice:   220, // 300 - 80
			wantApplied: 3,
		},
		{
			name:          "All discounts combined",
			originalPrice: 500,
			discounts: []DiscountInput{
				{Type: "fixed", Value: 50},
				{Type: "percentage", Value: 10},
				{Type: "conditional", Condition: 400, Value: 20},
				{
					Type: "tiered", Tiers: []Tier{
						{Min: 0, Max: 100, Value: 5},
						{Min: 101, Max: 499, Value: 15},
						{Min: 500, Max: 9999, Value: 30},
					},
				},
				{Type: "cap", MaxDiscount: 100},
			},
			wantPrice:   400,
			wantApplied: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrice, gotApplied := EvaluateDiscounts(tt.originalPrice, tt.discounts)

			if gotPrice != tt.wantPrice {
				t.Errorf("Got price = %v, want %v", gotPrice, tt.wantPrice)
			}

			if len(gotApplied) != tt.wantApplied {
				t.Errorf("Expected %d discounts applied, got %d", tt.wantApplied, len(gotApplied))
			}
		})
	}
}
