package carts

import "github.com/rovn208/df-go/ex06/internal/models/products"

type Cart struct {
	Products []products.Product `json:"products"`
}
