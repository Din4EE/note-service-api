package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	notes, err := n.repo.GetList(req.GetLimit(), req.GetOffset(), req.GetSearchQuery())
	if err != nil {
		return nil, err
	}

	var protoNotes []*desc.Note
	for _, note := range notes {
		protoNote := &desc.Note{
			Id:     note.ID,
			Title:  note.Title,
			Text:   note.Text,
			Author: note.Author,
		}
		protoNotes = append(protoNotes, protoNote)
	}

	totalCount := int64(len(notes))

	return &desc.GetListNoteResponse{
		Notes:      protoNotes,
		TotalCount: totalCount,
	}, nil
}
