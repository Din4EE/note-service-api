package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/app/repo"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {

	var noteUpdate repo.NoteUpdate
	if req.GetTitle() != nil {
		noteUpdate.Title = &req.GetTitle().Value
	}
	if req.GetText() != nil {
		noteUpdate.Text = &req.GetText().Value
	}
	if req.GetAuthor() != nil {
		noteUpdate.Author = &req.GetAuthor().Value
	}
	err := n.repo.Update(req.GetId(), noteUpdate)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
