package repo

import "context"

type Note struct {
	ID     string
	Title  string
	Text   string
	Author string
}

type NoteUpdate struct {
	Title  *string
	Text   *string
	Author *string
}

type NoteRepository interface {
	Create(ctx context.Context, note Note) (string, error)
	Get(ctx context.Context, id string) (Note, error)
	GetList(ctx context.Context, limit int64, offset int64, query string) ([]Note, error)
	Update(ctx context.Context, id string, note NoteUpdate) error
	Delete(ctx context.Context, id string) error
}
