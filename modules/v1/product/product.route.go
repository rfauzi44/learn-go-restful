package product

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()
	repo := NewRepo(db)
	controller := NewController(repo)

	route.HandleFunc("/", controller.Get).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")

}
