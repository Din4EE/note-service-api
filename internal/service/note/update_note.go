package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*emptypb.Empty, error) {
	note := converter.UpdateRequestToRepoNote(req)

	err := s.noteRepo.Update(ctx, req.GetId(), note)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
