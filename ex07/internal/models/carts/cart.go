package carts

type Cart struct {
	ProductCarts []ProductCart `json:"items"`
}

type ProductCart struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}
