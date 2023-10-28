package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := n.noteService.Create(ctx, converter.ToServiceNote(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
