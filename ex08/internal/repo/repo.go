package repo

import (
	"github.com/rovn208/df-go/ex08/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB(conn string) error {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderDetail{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
