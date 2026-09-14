package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SaH1l191/ticket-system/internal/model"
)

type PostgresTicketRepo struct {
	db *sql.DB
}

func NewPostgresTicketRepo(db *sql.DB) *PostgresTicketRepo {
	return &PostgresTicketRepo{db: db}
}

func (r *PostgresTicketRepo) Create(ctx context.Context, t model.Ticket) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO tickets (id, user_id, title, description, status, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		t.ID, t.UserID, t.Title, t.Description, t.Status, t.CreatedAt, t.UpdatedAt,
	)
	return err
}

func (r *PostgresTicketRepo) GetByID(ctx context.Context, id string) (*model.Ticket, error) {
	t := &model.Ticket{}
	err := r.db.QueryRowContext(ctx,
		`SELECT id, user_id, title, description, status, created_at, updated_at
		 FROM tickets WHERE id = $1`, id,
	).Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *PostgresTicketRepo) ListByUserID(ctx context.Context, userID string) ([]model.Ticket, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, user_id, title, description, status, created_at, updated_at
		 FROM tickets WHERE user_id = $1 ORDER BY created_at DESC`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []model.Ticket
	for rows.Next() {
		var t model.Ticket
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, rows.Err()
}

func (r *PostgresTicketRepo) UpdateStatus(ctx context.Context, id string, status string, now time.Time) error {
	result, err := r.db.ExecContext(ctx,
		"UPDATE tickets SET status = $1, updated_at = $2 WHERE id = $3",
		status, now, id,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("ticket not found")
	}
	return nil
}
