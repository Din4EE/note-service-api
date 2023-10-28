package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	notes, err := n.noteService.GetList(ctx, req.GetSearchQuery(), req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	protoNotes := make([]*desc.Note, 0, len(notes))
	for _, note := range notes {
		protoNotes = append(protoNotes, converter.ToDescNote(note))
	}

	return &desc.GetListResponse{
		Notes: protoNotes,
	}, nil
}
