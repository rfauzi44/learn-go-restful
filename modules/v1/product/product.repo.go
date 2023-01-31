package product

import (
	"github.com/rfauzi44/learn-go-restful/database/orm/model"
	"gorm.io/gorm"
)

type product_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}

}

func (r *product_repo) FindAll() (*model.Products, error) {
	var data model.Products
	err := r.database.Model(model.Product{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *product_repo) Add(data *model.Product) (*model.Product, error) {
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
