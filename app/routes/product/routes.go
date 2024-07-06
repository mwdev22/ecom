package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwdev22/ecom/app/utils"
)

type Handler struct {
	store ProductStore
}

func NewHandler(store ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

}

func (h *Handler) ProductList(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) NewProduct(w http.ResponseWriter, r *http.Request) {

}
