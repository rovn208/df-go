package model

type Product struct {
	Base
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=1"`
}

type ProductUri struct {
	ProductId string `uri:"product_id" binding:"required"`
}

type ProductIdJson struct {
	ProductId string `json:"product_id" binding:"required"`
}

type ProductItem struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}
