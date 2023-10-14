package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (s *Service) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	repoNote, err := s.noteRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	serviceNote := converter.RepoNoteToServiceNote(repoNote)

	return &desc.GetNoteResponse{
		Note: converter.ServiceNoteToDescNote(serviceNote),
	}, nil
}
