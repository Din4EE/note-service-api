package note

import (
	"context"
)

func (s *Service) Delete(ctx context.Context, id uint64) error {
	err := s.noteRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
