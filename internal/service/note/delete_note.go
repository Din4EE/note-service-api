package note

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := s.noteRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
