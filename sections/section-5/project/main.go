package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basedPrice, ok := productPrices[itemCode]
	if !ok {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItem := strings.TrimPrefix(itemCode, "_SALE")
			basedPrice, ok = productPrices[originalItem]
			if ok {
				salePrice := basedPrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItem, basedPrice, salePrice,
				)
				return salePrice, true
			}
		} else {
			fmt.Printf(" - Item %s not found\n", itemCode)
			return 0, false
		}
	}
	return basedPrice, true
}

func main() {
	fmt.Println("-------------- Simple Sales Order Processor------------")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subtotal float64
	fmt.Println("-------Processing Order Items:")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n", subtotal)
}
