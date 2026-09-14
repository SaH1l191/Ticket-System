package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaH1l191/ticket-system/internal/service"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, authResponse{Error: "invalid request body"})
		return
	}

	err := h.svc.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmailExists):
			writeJSON(w, http.StatusConflict, authResponse{Error: err.Error()})
		default:
			writeJSON(w, http.StatusBadRequest, authResponse{Error: err.Error()})
		}
		return
	}

	writeJSON(w, http.StatusCreated, authResponse{Message: "user created"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, authResponse{Error: "invalid request body"})
		return
	}

	token, err := h.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCredentials):
			writeJSON(w, http.StatusUnauthorized, authResponse{Error: err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, authResponse{Error: "internal error"})
		}
		return
	}

	writeJSON(w, http.StatusOK, authResponse{Token: token})
}
