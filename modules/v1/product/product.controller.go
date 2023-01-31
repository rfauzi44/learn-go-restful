package product

import (
	"encoding/json"
	"net/http"

	"github.com/rfauzi44/learn-go-restful/database/orm/model"
)

type product_controller struct {
	repo *product_repo
}

func NewController(repo *product_repo) *product_controller {
	return &product_controller{repo}
}

func (c *product_controller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response, err := c.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}

func (c *product_controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data model.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := c.repo.Add(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}
