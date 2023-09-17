package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	id := req.GetId()
	note, err := n.repo.Get(id)
	if err != nil {
		return nil, err
	}
	note.Title = req.GetTitle()
	note.Text = req.GetText()
	note.Author = req.GetAuthor()
	err = n.repo.Update(id, note)
	if err != nil {
		return nil, err
	}
	status := "updated"
	return &desc.UpdateNoteResponse{
		Id:     note.ID,
		Status: status,
	}, nil
}
