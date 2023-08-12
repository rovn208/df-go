package model

type Status string

var (
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

type Order struct {
	Base
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
	Status       Status        `json:"status"`
}

type OrderDetail struct {
	Base
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
