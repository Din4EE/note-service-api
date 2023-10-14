package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
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
		protoNote := converter.ServiceNoteToDescNote(serviceNote)
		protoNotes = append(protoNotes, protoNote)
	}

	return &desc.GetListNoteResponse{
		Notes: protoNotes,
	}, nil
}
