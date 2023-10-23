package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: converter.ToDescNote(note),
	}, nil
}
