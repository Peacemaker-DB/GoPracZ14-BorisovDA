package repo

import (
	"context"
	"time"

	"example.com/goprac11-borisovda/internal/core"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepoPG struct {
	pool *pgxpool.Pool
}

func NewNoteRepoPG(pool *pgxpool.Pool) *NoteRepoPG {
	return &NoteRepoPG{pool: pool}
}

func (r *NoteRepoPG) Create(ctx context.Context, n core.Note) (int64, error) {
	query := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id`
	var id int64
	err := r.pool.QueryRow(ctx, query, n.Title, n.Content).Scan(&id)
	return id, err
}

func (r *NoteRepoPG) GetAll(ctx context.Context, limit, offset int) ([]core.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes ORDER BY created_at DESC, id DESC LIMIT $1 OFFSET $2`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []core.Note{}
	for rows.Next() {
		var n core.Note
		var updatedAt *time.Time
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		n.UpdatedAt = updatedAt
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *NoteRepoPG) Get(ctx context.Context, id int64) (*core.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1`
	var n core.Note
	var updatedAt *time.Time
	err := r.pool.QueryRow(ctx, query, id).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	n.UpdatedAt = updatedAt
	return &n, nil
}

// GetBatch - батчинг для устранения N+1
func (r *NoteRepoPG) GetBatch(ctx context.Context, ids []int64) ([]core.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = ANY($1)`
	rows, err := r.pool.Query(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []core.Note{}
	for rows.Next() {
		var n core.Note
		var updatedAt *time.Time
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		n.UpdatedAt = updatedAt
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *NoteRepoPG) Update(ctx context.Context, id int64, upd core.Note) (*core.Note, error) {
	query := `UPDATE notes SET title = $1, content = $2, updated_at = now() WHERE id = $3 RETURNING id, title, content, created_at, updated_at`
	var n core.Note
	var updatedAt *time.Time
	err := r.pool.QueryRow(ctx, query, upd.Title, upd.Content, id).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	n.UpdatedAt = updatedAt
	return &n, nil
}

func (r *NoteRepoPG) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM notes WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}

// GetAllKeyset - keyset-пагинация вместо OFFSET
func (r *NoteRepoPG) GetAllKeyset(ctx context.Context, lastCreatedAt string, lastID int64, limit int) ([]core.Note, error) {
	query := `
		SELECT id, title, content, created_at, updated_at 
		FROM notes 
		WHERE (created_at, id) < ($1, $2)
		ORDER BY created_at DESC, id DESC 
		LIMIT $3`

	parsedTime, err := time.Parse(time.RFC3339, lastCreatedAt)
	if err != nil {
		parsedTime = time.Now()
	}

	rows, err := r.pool.Query(ctx, query, parsedTime, lastID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []core.Note{}
	for rows.Next() {
		var n core.Note
		var updatedAt *time.Time
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		n.UpdatedAt = updatedAt
		notes = append(notes, n)
	}
	return notes, nil
}
