package product

import "gorm.io/gorm"

type product_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}

}

func (r *product_repo) FindAll() string {
	return "hello from repo"
}
