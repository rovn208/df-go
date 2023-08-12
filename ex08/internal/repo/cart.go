package repo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rovn208/df-go/ex08/internal/constant"
	"github.com/rovn208/df-go/ex08/internal/model"
)

func AddProduct(pi model.ProductItem) error {
	p := &model.Product{}
	result := DB.First(&p, "id = ?", pi.ProductId)
	if result.Error != nil {
		return constant.InvalidProductIdError
	}
	order, err := getCurrentOrder()
	if err != nil {
		return err
	}
	for i, od := range order.OrderDetails {
		if od.ProductID == pi.ProductId {
			order.OrderDetails[i].Quantity += pi.Quantity
			return DB.Save(&order.OrderDetails[i]).Error
		}
	}

	order.OrderDetails = append(order.OrderDetails, model.OrderDetail{
		Base: model.Base{
			ID: uuid.New().String(),
		},
		ProductID: pi.ProductId,
		Quantity:  pi.Quantity,
		OrderID:   order.ID,
		Price:     p.Price,
	})

	return DB.Model(&order).Updates(order).Error
}

func Checkout() (string, error) {
	receipt := "Items:\n"
	total := 0.00
	order, err := getCurrentOrder()
	if err != nil {
		return "", err
	}
	for _, od := range order.OrderDetails {
		total += od.Price * float64(od.Quantity)
		receipt += fmt.Sprintf("ID: %s - Price: %0.2f - Quantity: %d\n", od.ProductID, od.Price, od.Quantity)
	}
	receipt += fmt.Sprintf("Total: %0.2f\n", total)
	order.Status = model.StatusCompleted
	return receipt, DB.Save(&order).Error
}

func RemoveItem(productId string) error {
	order, err := getCurrentOrder()
	if err != nil {
		return err
	}
	for i, od := range order.OrderDetails {
		if od.ProductID == productId {
			if od.Quantity == 1 {
				return DB.Delete(&order.OrderDetails[i]).Error
			}

			order.OrderDetails[i].Quantity -= 1
			return DB.Save(&order.OrderDetails[i]).Error
		}
	}

	return constant.ProductDoesNotExistsInCartError
}

// getCurrentOrder Gets the current order which has status is "in_progress" otherwise create a new one
func getCurrentOrder() (*model.Order, error) {
	order := &model.Order{}
	result := DB.Model(&order).Preload("OrderDetails").First(&order, "status = ?", model.StatusInProgress)
	if result.Error != nil {
		return createNewOrder()
	}
	return order, nil
}

func createNewOrder() (*model.Order, error) {
	order := &model.Order{
		Base: model.Base{
			ID: uuid.New().String(),
		},
		Status: model.StatusInProgress,
	}
	result := DB.Create(&order)
	if result.Error != nil {
		return order, result.Error
	}
	return order, nil
}
