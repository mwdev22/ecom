package product

import "github.com/gorilla/mux"

type Handler struct {
	store ProductStore
}

func NewHandler(store ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

}
