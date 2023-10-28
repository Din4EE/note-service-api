package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/model"
)

func (s *Service) Update(ctx context.Context, id uint64, note *model.NoteInfoUpdate) error {
	err := s.noteRepo.Update(ctx, id, note)
	if err != nil {
		return err
	}

	return nil
}
