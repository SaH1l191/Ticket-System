package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaH1l191/ticket-system/internal/service"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	svc *service.TicketService
}

func NewTicketHandler(s *service.TicketService) *TicketHandler {
	return &TicketHandler{svc: s}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)

	var req createTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ticketErrorResponse{Error: "invalid request body"})
		return
	}

	ticket, err := h.svc.Create(r.Context(), userID, req.Title, req.Description)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ticketErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, ticketResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *TicketHandler) ListTickets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)

	tickets, err := h.svc.ListByUser(r.Context(), userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ticketErrorResponse{Error: "could not list tickets"})
		return
	}

	resp := make([]ticketResponse, 0, len(tickets))
	for _, t := range tickets {
		resp = append(resp, ticketResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   t.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)
	ticketID := chi.URLParam(r, "id")

	ticket, err := h.svc.GetByID(r.Context(), userID, ticketID)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTicketNotFound), errors.Is(err, service.ErrAccessDenied):
			writeJSON(w, http.StatusNotFound, ticketErrorResponse{Error: "ticket not found"})
		default:
			writeJSON(w, http.StatusInternalServerError, ticketErrorResponse{Error: "internal error"})
		}
		return
	}

	writeJSON(w, http.StatusOK, ticketResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (h *TicketHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)
	ticketID := chi.URLParam(r, "id")

	var req updateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ticketErrorResponse{Error: "invalid request body"})
		return
	}

	ticket, err := h.svc.UpdateStatus(r.Context(), userID, ticketID, req.Status)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTicketNotFound), errors.Is(err, service.ErrAccessDenied):
			writeJSON(w, http.StatusNotFound, ticketErrorResponse{Error: "ticket not found"})
		case errors.Is(err, service.ErrInvalidStatus), errors.Is(err, service.ErrInvalidTransition):
			writeJSON(w, http.StatusBadRequest, ticketErrorResponse{Error: err.Error()})
		default:
			writeJSON(w, http.StatusInternalServerError, ticketErrorResponse{Error: "internal error"})
		}
		return
	}

	writeJSON(w, http.StatusOK, ticketResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      ticket.Status,
		CreatedAt:   ticket.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   ticket.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}
