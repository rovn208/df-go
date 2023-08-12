package repo

import (
	"github.com/rovn208/df-go/ex07/internal/constant"
	"github.com/rovn208/df-go/ex07/internal/model"
	"time"
)

func AddNewProduct(product model.Product) error {
	if DB.First(&product).Error == nil {
		return constant.ProductIsAlreadyExistsError
	}

	return DB.Create(&product).Error
}

func GetProducts() ([]model.Product, error) {
	var products []model.Product
	result := DB.Find(&products)
	return products, result.Error
}

func UpdateProduct(product model.Product) error {
	p := &model.Product{}
	if DB.First(&p, product.ID).Error != nil {
		return constant.InvalidProductIdError
	}
	p.Price = product.Price
	p.Description = product.Description
	p.UpdatedAt = time.Now()
	return DB.Save(&p).Error
}

func DeleteProduct(productId string) error {
	p := &model.Product{
		Base: model.Base{ID: productId},
	}
	if DB.First(&p).Error != nil {
		return constant.InvalidProductIdError
	}

	return DB.Delete(&p).Error
}
