package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/Din4EE/note-service-api/repo"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := n.repo.Update(req.GetId(), repo.Note{
		Title:  req.GetTitle().GetValue(),
		Author: req.GetAuthor().GetValue(),
		Text:   req.GetText().GetValue(),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
