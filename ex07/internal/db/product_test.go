package db

import (
	"github.com/rovn208/df-go/ex07/internal/constants"
	"github.com/rovn208/df-go/ex07/internal/models/products"
	"reflect"
	"testing"
)

func TestDB_AddNewProduct(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		Product       products.Product
		expectedDB    DB
		expectedError error
	}{
		{
			"add new product",
			DB{
				products: []products.Product{},
			},
			products.Product{ID: "1", Description: "", Price: 1},
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			nil,
		},
		{
			"product is already exists",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			products.Product{ID: "1", Description: "", Price: 1},
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			constants.ProductIsAlreadyExistsError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.DB.AddNewProduct(tc.Product)
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
		})
	}
}

func TestDB_GetProducts(t *testing.T) {
	testCases := []struct {
		name            string
		DB              DB
		expectedResults []products.Product
	}{
		{
			"get products",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			[]products.Product{{ID: "1", Description: "", Price: 1}},
		},
		{
			"products are empty",
			DB{
				products: []products.Product{},
			},
			[]products.Product{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := tc.DB.GetProducts()
			if !reflect.DeepEqual(results, tc.expectedResults) {
				t.Errorf("expected results: %v, got results: %v", tc.expectedResults, results)
			}
		})
	}
}

func TestDB_DeleteProduct(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		productId     string
		expectedDB    DB
		expectedError error
	}{
		{
			"delete product",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"1",
			DB{
				products: []products.Product{},
			},
			nil,
		},
		{
			"product is not exists if products are empty",
			DB{
				products: []products.Product{},
			},
			"1",
			DB{
				products: []products.Product{},
			},
			constants.InvalidProductIdError,
		},
		{
			"product is not exists",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"2",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			constants.InvalidProductIdError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.DB.DeleteProduct(tc.productId)
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
			if !reflect.DeepEqual(tc.DB.products, tc.expectedDB.products) {
				t.Errorf("expected products: %v, got products: %v", tc.expectedDB.products, tc.DB.products)
			}
		})
	}
}

func TestDB_GetProductById(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		productId     string
		expectedDB    products.Product
		expectedError error
	}{
		{
			"get product by id",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"1",
			products.Product{ID: "1", Description: "", Price: 1},
			nil,
		},
		{
			"product is not exists",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			"2",
			products.Product{},
			constants.InvalidProductIdError,
		},
		{
			"product is not exists and products are empty",
			DB{
				products: []products.Product{},
			},
			"2",
			products.Product{},
			constants.InvalidProductIdError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			product, err := tc.DB.GetProductById(tc.productId)
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
			if !reflect.DeepEqual(product, tc.expectedDB) {
				t.Errorf("expected product: %v, got product: %v", tc.expectedDB, product)
			}
		})
	}
}

func TestDB_UpdateProduct(t *testing.T) {
	testCases := []struct {
		name          string
		DB            DB
		Product       products.Product
		expectedDB    DB
		expectedError error
	}{
		{
			"update product",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			products.Product{ID: "1", Description: "", Price: 2},
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 2}},
			},
			nil,
		},
		{
			"product is not exists",
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			products.Product{ID: "2", Description: "", Price: 2},
			DB{
				products: []products.Product{{ID: "1", Description: "", Price: 1}},
			},
			constants.InvalidProductIdError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.DB.UpdateProduct(tc.Product)
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
			if !reflect.DeepEqual(tc.DB.products, tc.expectedDB.products) {
				t.Errorf("expected products: %v, got products: %v", tc.expectedDB.products, tc.DB.products)
			}
		})
	}
}
