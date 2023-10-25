package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(ctx, req.GetId(), converter.ToServiceNoteFromUpdate(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
