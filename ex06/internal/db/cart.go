package db

import (
	"fmt"
	"github.com/rovn208/df-go/ex06/internal/constants"
	"github.com/rovn208/df-go/ex06/internal/models/carts"
	"github.com/rovn208/df-go/ex06/internal/util"
)

func (db *DB) AddProduct(pc carts.ProductCart) error {
	if _, err := db.GetProductById(pc.ProductId); err != nil {
		return constants.InvalidProductIdError
	}

	for _, productCart := range db.cart.ProductCarts {
		if productCart.ProductId == pc.ProductId {
			productCart.Quantity += pc.Quantity
			return nil
		}
	}

	db.cart.ProductCarts = append(db.cart.ProductCarts, pc)
	return nil
}

func (db *DB) Checkout() (string, error) {
	if len(db.cart.ProductCarts) == 0 {
		return "Total: 0.00", nil
	}

	total := 0.00
	receipt := "Items:\n"
	for _, pc := range db.cart.ProductCarts {
		product, err := db.GetProductById(pc.ProductId)
		if err != nil {
			return "", err
		}
		total += float64(pc.Quantity) * product.Price
		receipt += fmt.Sprintf("ID: %s - Price: %0.2f - Quantity: %d\n", pc.ProductId, product.Price, pc.Quantity)
	}
	receipt += fmt.Sprintf("Total: %0.2f\n", total)
	db.cart.ProductCarts = make([]carts.ProductCart, 0)
	return receipt, nil
}

func (db *DB) RemoveItem(productId string) error {
	for idx, productCart := range db.cart.ProductCarts {
		if productCart.ProductId == productId {
			if productCart.Quantity == 1 {
				pcs, err := util.RemoveProductCartAt(db.cart.ProductCarts, idx)
				if err != nil {
					return err
				}
				db.cart.ProductCarts = pcs
				return nil
			}
			db.cart.ProductCarts[idx].Quantity--
			return nil
		}
	}

	return constants.ProductDoesNotExistsInCartError
}
