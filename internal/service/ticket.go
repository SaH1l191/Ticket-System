package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/SaH1l191/ticket-system/internal/model"
	"github.com/SaH1l191/ticket-system/internal/repo"
	"github.com/google/uuid"
)

var (
	ErrTicketNotFound       = errors.New("ticket not found")
	ErrAccessDenied         = errors.New("ticket not found")
	ErrInvalidStatus        = errors.New("invalid status. allowed: open, in_progress, closed")
	ErrInvalidTransition    = errors.New("invalid status transition")
)

var validStatuses = map[string]bool{
	"open":        true,
	"in_progress": true,
	"closed":      true,
}

var allowedTransitions = map[string]string{
	"open":        "in_progress",
	"in_progress": "closed",
}

type TicketService struct {
	tickets repo.TicketRepository
}

func NewTicketService(t repo.TicketRepository) *TicketService {
	return &TicketService{tickets: t}
}

func (s *TicketService) Create(ctx context.Context, userID, title, description string) (*model.Ticket, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("title is required")
	}

	now := time.Now()
	ticket := model.Ticket{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      "open",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.tickets.Create(ctx, ticket); err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (s *TicketService) ListByUser(ctx context.Context, userID string) ([]model.Ticket, error) {
	return s.tickets.ListByUserID(ctx, userID)
}

func (s *TicketService) GetByID(ctx context.Context, userID, ticketID string) (*model.Ticket, error) {
	ticket, err := s.tickets.GetByID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, ErrTicketNotFound
	}
	if ticket.UserID != userID {
		return nil, ErrAccessDenied
	}
	return ticket, nil
}

func (s *TicketService) UpdateStatus(ctx context.Context, userID, ticketID, newStatus string) (*model.Ticket, error) {
	ticket, err := s.tickets.GetByID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, ErrTicketNotFound
	}
	if ticket.UserID != userID {
		return nil, ErrAccessDenied
	}

	newStatus = strings.TrimSpace(newStatus)
	if !validStatuses[newStatus] {
		return nil, ErrInvalidStatus
	}

	expectedNext, ok := allowedTransitions[ticket.Status]
	if !ok || expectedNext != newStatus {
		return nil, ErrInvalidTransition
	}

	if err := s.tickets.UpdateStatus(ctx, ticketID, newStatus, time.Now()); err != nil {
		return nil, err
	}

	return s.tickets.GetByID(ctx, ticketID)
}
