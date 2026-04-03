package post

import (
	"context"

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

func (p *Pgx) CreatePost(ctx context.Context, post *model.APost) (string, error) {
	query := `
    INSERT INTO animals (
      organization_id, 
      name, 
      age, 
      sex, 
      description, 
      photo_url, 
      status
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id
  `

	var id string
	err := p.pool.QueryRow(
		ctx,
		query,
		post.OrganizationID,
		post.Name,
		post.Age,
		post.Sex,
		post.Description,
		post.PhotoURL,
		post.Status,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
