package db

import (
	"github.com/rovn208/df-go/ex07/internal/constants"
	"github.com/rovn208/df-go/ex07/internal/models/carts"
	"github.com/rovn208/df-go/ex07/internal/models/products"
	"reflect"
	"testing"
)

func TestDB_AddProduct(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		pc            carts.ProductCart
		expectedDB    DB
		expectedError error
	}{
		{
			"add new product",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			carts.ProductCart{ProductId: "1", Quantity: 1},
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			nil,
		},
		{
			"update quantity if product is already exists in cart",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			carts.ProductCart{ProductId: "1", Quantity: 2},
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 3}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			nil,
		},
		{
			"product does not exists",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			carts.ProductCart{ProductId: "2", Quantity: 1},
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			constants.InvalidProductIdError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.DB.AddProduct(tc.pc)
			if err != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
					return
				}
			} else {
				if tc.expectedError != nil {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
				}
			}
			if !reflect.DeepEqual(tc.expectedDB, tc.DB) {
				t.Errorf("expected db: %v, got db: %v", tc.expectedDB, tc.DB)
			}
		})
	}
}

func TestDB_Checkout(t *testing.T) {
	testCases := []struct {
		name            string
		DB              DB
		expectedReceipt string
		expectedError   error
	}{
		{
			"checkout",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"Items:\nID: 1 - Price: 1.00 - Quantity: 1\nTotal: 1.00\n",
			nil,
		},
		{
			"cart is empty",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"Total: 0.00",
			nil,
		},
		{
			"product cart item is not valid",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "2", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"",
			constants.InvalidProductIdError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			receipt, err := tc.DB.Checkout()
			if err != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
					return
				}
			} else {
				if tc.expectedError != nil {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
				}
			}
			if !reflect.DeepEqual(receipt, tc.expectedReceipt) {
				t.Errorf("expected receipt: %s, got receipt: %s", tc.expectedReceipt, receipt)
			}
		})
	}
}

func TestDB_RemoveItem(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		productId     string
		expectedDB    DB
		expectedError error
	}{
		{
			"remove item when quantity is equal 1",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"1",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			nil,
		},
		{
			"remove item when quantity is greater than 1",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 2}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"1",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			nil,
		},
		{
			"product does not exists",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"2",
			DB{
				cart: carts.Cart{
					ProductCarts: []carts.ProductCart{{ProductId: "1", Quantity: 1}},
				},
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			constants.ProductDoesNotExistsInCartError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.DB.RemoveItem(tc.productId)
			if err != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
					return
				}
			} else {
				if tc.expectedError != nil {
					t.Errorf("expected error: %s, got error: %s", tc.expectedError.Error(), err.Error())
				}
			}
			if !reflect.DeepEqual(tc.expectedDB, tc.DB) {
				t.Errorf("expected db: %v, got db: %v", tc.expectedDB, tc.DB)
			}
		})
	}
}
