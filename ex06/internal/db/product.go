package db

import (
	"github.com/rovn208/df-go/ex06/internal/constants"
	"github.com/rovn208/df-go/ex06/internal/models/products"
	"github.com/rovn208/df-go/ex06/internal/util"
)

func (db *DB) AddNewProduct(product products.Product) error {
	if util.IsProductExists(db.products, product) {
		return constants.ProductIsAlreadyExistsError
	}

	db.products = append(db.products, product)
	return nil
}

func (db *DB) GetProducts() []products.Product {
	return db.products
}

func (db *DB) UpdateProduct(product products.Product) error {
	for idx, p := range db.products {
		if p.ID == product.ID {
			db.products[idx] = product
			return nil
		}
	}
	return nil
}

func (db *DB) DeleteProduct(productId string) error {
	for idx, p := range db.products {
		if p.ID == productId {
			db.products = append(db.products[:idx], db.products[idx+1:]...)
			return nil
		}
	}

	return constants.InvalidProductIdError
}

func (db *DB) GetProductById(productId string) (products.Product, error) {
	for _, p := range db.products {
		if p.ID == productId {
			return p, nil
		}
	}

	return products.Product{}, constants.InvalidProductIdError
}
