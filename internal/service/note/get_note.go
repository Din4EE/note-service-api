package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	repoNote, err := s.noteRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	serviceNote := converter.RepoNoteToServiceNote(repoNote)

	return &desc.GetNoteResponse{
		Note: &desc.Note{
			Id:        serviceNote.ID,
			Title:     serviceNote.Title,
			Text:      serviceNote.Text,
			Author:    serviceNote.Author,
			Email:     serviceNote.Email,
			CreatedAt: timestamppb.New(serviceNote.CreatedAt),
			UpdatedAt: timestamppb.New(serviceNote.UpdatedAt),
		},
	}, nil
}
