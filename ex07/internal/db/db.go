package db

import (
	"github.com/rovn208/df-go/ex07/internal/models/carts"
	"github.com/rovn208/df-go/ex07/internal/models/products"
)

// DB is a mock DB which contains a local cart and also provide some mock functions to interact with this DB
type DB struct {
	products []products.Product
	cart     carts.Cart
}

var MockDB *DB

func init() {
	MockDB = &DB{
		products: make([]products.Product, 0),
		cart:     carts.Cart{ProductCarts: make([]carts.ProductCart, 0)}}
}
