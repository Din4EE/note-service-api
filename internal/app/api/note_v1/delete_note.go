package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (empty *emptypb.Empty, err error) {
	id := req.GetId()
	err = n.repo.Delete(id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
