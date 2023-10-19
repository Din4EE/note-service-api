package note_v1

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(ctx, converter.DescNoteToServiceNote(&desc.Note{
		Id:     req.GetId(),
		Title:  req.GetTitle().GetValue(),
		Text:   req.GetText().GetValue(),
		Author: req.GetAuthor().GetValue(),
		Email:  req.GetEmail().GetValue(),
	}))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
