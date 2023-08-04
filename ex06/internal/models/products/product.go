package products

type Product struct {
	ID          string  `json:"id" binding:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" binding:"required"`
}
