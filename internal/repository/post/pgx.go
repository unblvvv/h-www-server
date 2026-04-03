package post

import (
	"context"
	"errors"

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

func (p *Pgx) DeletePost(ctx context.Context, id, userId string) error {
	query := `DELETE FROM animals WHERE id = $1 AND organization_id = $2`

	tag, err := p.pool.Exec(ctx, query, id, userId)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("post not found or user not authorized to delete")
	}

	return nil
}

func (p *Pgx) GetPost(ctx context.Context, limit, offset int) ([]model.APost, error) {
	query := `
        SELECT id, organization_id, name, age, sex, description, photo_url, status, created_at, updated_at
        FROM animals
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
	rows, err := p.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []model.APost
	for rows.Next() {
		var a model.APost
		err := rows.Scan(&a.ID, &a.OrganizationID, &a.Name, &a.Age, &a.Sex, &a.Description, &a.PhotoURL, &a.Status, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		animals = append(animals, a)
	}
	return animals, rows.Err()
}

func (p *Pgx) UpdatePost(ctx context.Context, name, age, sex, description string, photo_url *string, post_id, userId string) error {
	query := `
       UPDATE animals 
       SET name = $1, age = $2, sex = $3, description = $4, photo_url = $5, updated_at = NOW()
       WHERE id = $6 AND organization_id = $7
    `

	tag, err := p.pool.Exec(ctx, query, name, age, sex, description, photo_url, post_id, userId)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("post not found or user not authorized to delete")
	}

	return nil
}
