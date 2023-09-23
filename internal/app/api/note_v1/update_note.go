package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	note, err := n.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if title := req.GetTitle(); title != nil {
		note.Title = title.GetValue()
	}
	if text := req.GetText(); text != nil {
		note.Text = text.GetValue()
	}
	if author := req.GetAuthor(); author != nil {
		note.Author = author.GetValue()
	}

	err = n.repo.Update(id, note)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
