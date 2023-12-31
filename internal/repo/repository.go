package repo

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_note_repository.go -package=mocks . NoteRepository

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/model"
)

type NoteRepository interface {
	Create(ctx context.Context, note *model.Note) (uint64, error)
	Get(ctx context.Context, id uint64) (*model.Note, error)
	GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error)
	Update(ctx context.Context, id uint64, note *model.NoteInfoUpdate) error
	Delete(ctx context.Context, id uint64) error
}
