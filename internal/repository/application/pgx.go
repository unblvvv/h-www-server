package application

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unblvvv/h-www-server/internal/model"
)

type Pgx struct {
	pool *pgxpool.Pool
}

func NewFx(pool *pgxpool.Pool) Repository {
	return &Pgx{pool: pool}
}

func (r *Pgx) Create(ctx context.Context, app *model.Application) error {
	query := `INSERT INTO applications (animal_id, name, email, phone, message) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.pool.Exec(ctx, query, app.AnimalID, app.Name, app.Email, app.Phone, app.Message)
	return err
}

func (r *Pgx) GetList(ctx context.Context, status *string, limit, offset int) ([]model.Application, int, error) {
	var items []model.Application
	var total int

	countQuery := `SELECT COUNT(*) FROM applications WHERE ($1::text IS NULL OR status = $1)`
	err := r.pool.QueryRow(ctx, countQuery, status).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			a.id, a.animal_id, an.name as animal_name, 
			a.name, a.email, a.phone, a.message, a.status, a.created_at 
		FROM applications a
		JOIN animals an ON a.animal_id = an.id
		WHERE ($1::text IS NULL OR a.status = $1)
		ORDER BY a.created_at DESC 
		LIMIT $2 OFFSET $3`

	rows, err := r.pool.Query(ctx, query, status, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Application
		err := rows.Scan(&a.ID, &a.AnimalID, &a.AnimalName, &a.Name, &a.Email, &a.Phone, &a.Message, &a.Status, &a.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, a)
	}

	return items, total, nil
}

func (r *Pgx) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE applications SET status = $1 WHERE id = $2`
	_, err := r.pool.Exec(ctx, query, status, id)
	return err
}
