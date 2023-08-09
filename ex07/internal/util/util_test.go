package util

import (
	"github.com/rovn208/df-go/ex07/internal/constants"
	"github.com/rovn208/df-go/ex07/internal/models/carts"
	"github.com/rovn208/df-go/ex07/internal/models/products"
	"testing"
)

func TestIsProductExists(t *testing.T) {
	testCases := []struct {
		name     string
		products []products.Product
		product  products.Product
		expected bool
	}{
		{"product exists",
			[]products.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
			products.Product{ID: "1", Description: "", Price: 1},
			true,
		},
		{"products are empty",
			[]products.Product{},
			products.Product{ID: "1", Description: "", Price: 1},
			false,
		},
		{"Product does not exist}",
			[]products.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
			products.Product{ID: "3", Description: "", Price: 1},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := IsProductExists(tc.products, tc.product); actual != tc.expected {
				t.Errorf("IsProductExists() = %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestRemoveProductCartAt(t *testing.T) {
	testCases := []struct {
		name          string
		productCards  []carts.ProductCart
		index         int
		expected      []carts.ProductCart
		expectedError error
	}{
		{
			"remove product cart at index 0",
			[]carts.ProductCart{{ProductId: "1", Quantity: 1}},
			0,
			[]carts.ProductCart{},
			nil,
		},
		{
			"remove product cart at index 1",
			[]carts.ProductCart{{ProductId: "1", Quantity: 1}, {ProductId: "2", Quantity: 1}},
			1,
			[]carts.ProductCart{{ProductId: "1", Quantity: 1}},
			nil,
		},
		{
			"index is greater than length of product cart",
			[]carts.ProductCart{{ProductId: "1", Quantity: 1}},
			2,
			[]carts.ProductCart{},
			constants.InvalidIndexError,
		},
		{
			"product carts are empty",
			[]carts.ProductCart{},
			0,
			[]carts.ProductCart{},
			constants.InvalidIndexError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RemoveProductCartAt(tc.productCards, tc.index)
			if len(actual) != len(tc.expected) {
				t.Errorf("RemoveProductCartAt() = %v, want %v", actual, tc.expected)
			}
			for i, productCart := range actual {
				if productCart.ProductId != tc.expected[i].ProductId {
					t.Errorf("RemoveProductCartAt() = %v, want %v", actual, tc.expected)
				}
			}
			if err != tc.expectedError {
				t.Errorf("RemoveProductCartAt() = %v, want %v", err, tc.expectedError)
			}
		})
	}
}

func TestRemoveProductAt(t *testing.T) {
	testCases := []struct {
		name          string
		productList   []products.Product
		index         int
		expected      []products.Product
		expectedError error
	}{
		{
			"remove product at index 0",
			[]products.Product{{ID: "1", Description: "", Price: 1}},
			0,
			[]products.Product{},
			nil,
		},
		{
			"remove product at index 1",
			[]products.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
			1,
			[]products.Product{{ID: "1", Description: "", Price: 1}},
			nil,
		},
		{
			"index is greater than length of product list",
			[]products.Product{{ID: "1", Description: "", Price: 1}},
			2,
			[]products.Product{},
			constants.InvalidIndexError,
		},
		{
			"product list is empty",
			[]products.Product{},
			0,
			[]products.Product{},
			constants.InvalidIndexError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := RemoveProductAt(tc.productList, tc.index)
			if len(actual) != len(tc.expected) {
				t.Errorf("RemoveProductAt() = %v, want %v", actual, tc.expected)
			}
			for i, product := range actual {
				if product.ID != tc.expected[i].ID {
					t.Errorf("RemoveProductAt() = %v, want %v", actual, tc.expected)
				}
			}
			if err != tc.expectedError {
				t.Errorf("RemoveProductAt() = %v, want %v", err, tc.expectedError)
			}
		})
	}
}
