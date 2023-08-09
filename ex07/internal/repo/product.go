package repo

import (
	"github.com/rovn208/df-go/ex07/internal/constant"
	"github.com/rovn208/df-go/ex07/internal/model"
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
	if DB.First(&product).Error != nil {
		return constant.InvalidProductIdError
	}
	DB.Save(&product)
	return nil
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
