package router

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/learn-go-restful/database/orm"
	"github.com/rfauzi44/learn-go-restful/modules/v1/product"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.NewDB()
	if err != nil {
		return nil, err
	}

	product.New(mainRoute, db)
	return mainRoute, nil
}
