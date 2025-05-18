package main

func EvaluateDiscounts(originalPrice float64, discounts []DiscountInput) (finalPrice float64, applied []map[string]interface{}) {
	price := originalPrice
	totalDiscount := 0.0
	var capLimit float64 = -1

	for _, d := range discounts {
		switch d.Type {
		case "fixed":
			amount := d.Value
			price -= amount
			totalDiscount += amount
			applied = append(applied, map[string]interface{}{"type": "fixed", "amount": amount})
		case "percentage":
			amount := price * d.Value / 100
			price -= amount
			totalDiscount += amount
			applied = append(applied, map[string]interface{}{"type": "percentage", "amount": amount})
		case "conditional":
			if originalPrice > d.Condition {
				amount := d.Value
				price -= amount
				totalDiscount += amount
				applied = append(applied, map[string]interface{}{"type": "conditional", "amount": amount})
			}
		case "tiered":
			for _, tier := range d.Tiers {
				if originalPrice >= tier.Min && originalPrice <= tier.Max {
					amount := tier.Value
					price -= amount
					totalDiscount += amount
					applied = append(applied, map[string]interface{}{"type": "tiered", "amount": amount})
					break
				}
			}
		case "cap":
			capLimit = d.MaxDiscount
		}
	}

	// Apply cap at the end
	if capLimit >= 0 && totalDiscount > capLimit {
		price = originalPrice - capLimit
		applied = append(applied, map[string]interface{}{
			"type":                  "cap",
			"originalDiscountTotal": totalDiscount,
			"cappedAt":              capLimit,
		})
	}

	if price < 0 {
		price = 0
	}

	return price, applied
}
