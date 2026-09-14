package repo

import (
	"context"
	"database/sql"

	"github.com/SaH1l191/ticket-system/internal/model"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(ctx context.Context, u model.User) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (id, email, password_hash, created_at) VALUES ($1, $2, $3, $4)",
		u.ID, u.Email, u.PasswordHash, u.CreatedAt,
	)
	return err
}

func (r *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	u := &model.User{}
	err := r.db.QueryRowContext(ctx,
		"SELECT id, email, password_hash, created_at FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
