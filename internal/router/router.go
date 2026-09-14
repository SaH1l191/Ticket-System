package router

import (
	"encoding/json"
	"net/http"

	"github.com/SaH1l191/ticket-system/internal/handler"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func New(authH *handler.AuthHandler, ticketH *handler.TicketHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RequestID)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authH.Register)
		r.Post("/login", authH.Login)
	})

	r.Route("/tickets", func(r chi.Router) {
		r.Use(handler.AuthMiddleware)
		r.Post("/", ticketH.CreateTicket)
		r.Get("/", ticketH.ListTickets)
		r.Get("/{id}", ticketH.GetTicket)
		r.Patch("/{id}/status", ticketH.UpdateStatus)
	})

	return r
}
