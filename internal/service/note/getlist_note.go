package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	if req.GetLimit() == 0 {
		req.Limit = 150
	}

	repoNotes, err := s.noteRepo.GetList(ctx, req.GetSearchQuery(), req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	var protoNotes []*desc.Note
	for _, repoNote := range repoNotes {
		serviceNote := converter.RepoNoteToServiceNote(repoNote)

		protoNote := &desc.Note{
			Id:        serviceNote.ID,
			Title:     serviceNote.Title,
			Text:      serviceNote.Text,
			Author:    serviceNote.Author,
			Email:     serviceNote.Email,
			CreatedAt: timestamppb.New(serviceNote.CreatedAt),
			UpdatedAt: timestamppb.New(serviceNote.UpdatedAt),
		}
		protoNotes = append(protoNotes, protoNote)
	}

	return &desc.GetListNoteResponse{
		Notes: protoNotes,
	}, nil
}
