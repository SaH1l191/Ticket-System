package repo

import (
	"context"
	"time"

	"github.com/SaH1l191/ticket-system/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, u model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type TicketRepository interface {
	Create(ctx context.Context, t model.Ticket) error
	GetByID(ctx context.Context, id string) (*model.Ticket, error)
	ListByUserID(ctx context.Context, userID string) ([]model.Ticket, error)
	UpdateStatus(ctx context.Context, id string, status string, now time.Time) error
}
