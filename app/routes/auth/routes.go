package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwdev22/ecom/app/types"
	"github.com/mwdev22/ecom/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/register", h.Register).Methods("POST")
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err) // bad payload
		return
	}
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password"))
		return
	}

	token, err := GenerateJWT(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateUser(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	msg := make(map[string]string)
	msg["success"] = "user created successfully"
	utils.WriteJSON(w, http.StatusCreated, msg)
}
