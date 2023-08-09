package util

//func TestIsProductExists(t *testing.T) {
//	testCases := []struct {
//		name     string
//		products []model.Product
//		product  model.Product
//		expected bool
//	}{
//		{"product exists",
//			[]model.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
//			model.Product{ID: "1", Description: "", Price: 1},
//			true,
//		},
//		{"products are empty",
//			[]model.Product{},
//			model.Product{ID: "1", Description: "", Price: 1},
//			false,
//		},
//		{"Product does not exist}",
//			[]model.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
//			model.Product{ID: "3", Description: "", Price: 1},
//			false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			if actual := IsProductExists(tc.products, tc.product); actual != tc.expected {
//				t.Errorf("IsProductExists() = %v, want %v", actual, tc.expected)
//			}
//		})
//	}
//}
//
//func TestRemoveProductCartAt(t *testing.T) {
//	testCases := []struct {
//		name          string
//		productCards  []model.ProductCart
//		index         int
//		expected      []model.ProductCart
//		expectedError error
//	}{
//		{
//			"remove product cart at index 0",
//			[]model.ProductCart{{ProductId: "1", Quantity: 1}},
//			0,
//			[]model.ProductCart{},
//			nil,
//		},
//		{
//			"remove product cart at index 1",
//			[]model.ProductCart{{ProductId: "1", Quantity: 1}, {ProductId: "2", Quantity: 1}},
//			1,
//			[]model.ProductCart{{ProductId: "1", Quantity: 1}},
//			nil,
//		},
//		{
//			"index is greater than length of product cart",
//			[]model.ProductCart{{ProductId: "1", Quantity: 1}},
//			2,
//			[]model.ProductCart{},
//			constant.InvalidIndexError,
//		},
//		{
//			"product carts are empty",
//			[]model.ProductCart{},
//			0,
//			[]model.ProductCart{},
//			constant.InvalidIndexError,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			actual, err := RemoveProductCartAt(tc.productCards, tc.index)
//			if len(actual) != len(tc.expected) {
//				t.Errorf("RemoveProductCartAt() = %v, want %v", actual, tc.expected)
//			}
//			for i, productCart := range actual {
//				if productCart.ProductId != tc.expected[i].ProductId {
//					t.Errorf("RemoveProductCartAt() = %v, want %v", actual, tc.expected)
//				}
//			}
//			if err != tc.expectedError {
//				t.Errorf("RemoveProductCartAt() = %v, want %v", err, tc.expectedError)
//			}
//		})
//	}
//}
//
//func TestRemoveProductAt(t *testing.T) {
//	testCases := []struct {
//		name          string
//		productList   []model.Product
//		index         int
//		expected      []model.Product
//		expectedError error
//	}{
//		{
//			"remove product at index 0",
//			[]model.Product{{ID: "1", Description: "", Price: 1}},
//			0,
//			[]model.Product{},
//			nil,
//		},
//		{
//			"remove product at index 1",
//			[]model.Product{{ID: "1", Description: "", Price: 1}, {ID: "2", Description: "", Price: 2}},
//			1,
//			[]model.Product{{ID: "1", Description: "", Price: 1}},
//			nil,
//		},
//		{
//			"index is greater than length of product list",
//			[]model.Product{{ID: "1", Description: "", Price: 1}},
//			2,
//			[]model.Product{},
//			constant.InvalidIndexError,
//		},
//		{
//			"product list is empty",
//			[]model.Product{},
//			0,
//			[]model.Product{},
//			constant.InvalidIndexError,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			actual, err := RemoveProductAt(tc.productList, tc.index)
//			if len(actual) != len(tc.expected) {
//				t.Errorf("RemoveProductAt() = %v, want %v", actual, tc.expected)
//			}
//			for i, product := range actual {
//				if product.ID != tc.expected[i].ID {
//					t.Errorf("RemoveProductAt() = %v, want %v", actual, tc.expected)
//				}
//			}
//			if err != tc.expectedError {
//				t.Errorf("RemoveProductAt() = %v, want %v", err, tc.expectedError)
//			}
//		})
//	}
//}
