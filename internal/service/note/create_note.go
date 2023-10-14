package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/converter"
	"github.com/Din4EE/note-service-api/internal/service"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (s *Service) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	repoNote := converter.ServiceNoteToRepoNote(&service.Note{
		Title:  req.GetTitle(),
		Text:   req.GetText(),
		Email:  req.GetEmail(),
		Author: req.GetAuthor(),
	},
	)

	id, err := s.noteRepo.Create(ctx, repoNote)
	if err != nil {
		return nil, err
	}

	return &desc.CreateNoteResponse{
		Id: id,
	}, nil
}
