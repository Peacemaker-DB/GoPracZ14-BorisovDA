package repo

import (
	"context"

	"example.com/goprac11-borisovda/internal/core"
)

type NoteRepository interface {
	Create(ctx context.Context, n core.Note) (int64, error)
	GetAll(ctx context.Context, limit, offset int) ([]core.Note, error)
	Get(ctx context.Context, id int64) (*core.Note, error)
	GetBatch(ctx context.Context, ids []int64) ([]core.Note, error)
	Update(ctx context.Context, id int64, upd core.Note) (*core.Note, error)
	Delete(ctx context.Context, id int64) error
	GetAllKeyset(ctx context.Context, lastCreatedAt string, lastID int64, limit int) ([]core.Note, error)
}
