package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/Din4EE/note-service-api/repo"
)

func (n *Note) CreateNote(ctx context.Context, res *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	newNote := repo.Note{
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Author: res.GetAuthor(),
	}
	id, err := n.repo.Create(newNote)
	if err != nil {
		return nil, err
	}
	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
