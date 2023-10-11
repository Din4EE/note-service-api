package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/app/repo"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, res *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	id, err := n.repo.Create(ctx, repo.Note{
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Author: res.GetAuthor(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
