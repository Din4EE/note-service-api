package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/model"
)

func (s *Service) Update(ctx context.Context, note *model.Note) error {
	err := s.noteRepo.Update(ctx, note)
	if err != nil {
		return err
	}

	return nil
}
