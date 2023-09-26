package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	note, err := n.repo.Get(req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetNoteResponse{
		Note: &desc.Note{
			Id:     note.ID,
			Title:  note.Title,
			Text:   note.Text,
			Author: note.Author,
		},
	}, nil
}
