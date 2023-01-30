package product

import "net/http"

type product_controller struct {
	repo *product_repo
}

func NewController(repo *product_repo) *product_controller {
	return &product_controller{repo}
}

func (c *product_controller) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response := c.repo.FindAll()
	w.Write([]byte(response))

}
