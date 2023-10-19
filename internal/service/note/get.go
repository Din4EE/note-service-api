package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/repo/converter"
	"github.com/Din4EE/note-service-api/internal/service/model"
)

func (s *Service) Get(ctx context.Context, id uint64) (*model.Note, error) {
	repoNote, err := s.noteRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.RepoNoteToServiceNote(repoNote), nil
}
