package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/service/model"
)

const (
	defaultLimit = 150
)

func (s *Service) GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error) {
	if limit == 0 {
		limit = defaultLimit
	}

	notes, err := s.noteRepo.GetList(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
