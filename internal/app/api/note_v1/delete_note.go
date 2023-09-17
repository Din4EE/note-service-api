package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (res *desc.DeleteNoteResponse, err error) {
	id := req.GetId()
	err = n.repo.Delete(id)
	if err != nil {
		return nil, err
	}
	status := "deleted"
	return &desc.DeleteNoteResponse{Status: status}, nil
}
