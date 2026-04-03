package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unblvvv/h-www-server/internal/model"
)

type Opts struct {
	PgxPool *pgxpool.Pool
}

type Pgx struct {
	pool *pgxpool.Pool
}

func New(opts Opts) Repository {
	return &Pgx{
		pool: opts.PgxPool,
	}
}

func NewFx(pool *pgxpool.Pool) Repository {
	return New(Opts{
		PgxPool: pool,
	})
}

func (p *Pgx) CreateUser(ctx context.Context, user model.User) (string, error) {
	query := `
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id string
	err := p.pool.QueryRow(ctx, query, user.Email, user.Username, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (p *Pgx) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
       SELECT id, email, username, password_hash, role
       FROM users WHERE email = $1
    `
	var user model.User

	err := p.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
