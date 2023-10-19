package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/repo/converter"
	"github.com/Din4EE/note-service-api/internal/service/model"
)

func (s *Service) Create(ctx context.Context, note *model.Note) (uint64, error) {
	id, err := s.noteRepo.Create(ctx, converter.ServiceNoteToRepoNote(note))
	if err != nil {
		return 0, err
	}

	return id, nil
}
