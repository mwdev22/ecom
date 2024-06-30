package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwdev22/ecom/app/types"
	"github.com/mwdev22/ecom/app/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/register", h.Register).Methods("POST")
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err) // bad request, problem on the user's
	}

}
