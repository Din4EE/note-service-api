package repo

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/repo/model"
)

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) (uint64, error)
	Get(ctx context.Context, id uint64) (*model.Note, error)
	GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error)
	Update(ctx context.Context, note *model.Note) error
	Delete(ctx context.Context, id uint64) error
}
