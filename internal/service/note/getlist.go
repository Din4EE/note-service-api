package note

import (
	"context"

	"github.com/Din4EE/note-service-api/internal/repo/converter"
	"github.com/Din4EE/note-service-api/internal/service/model"
)

const (
	defaultLimit = 150
)

func (s *Service) GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error) {
	if limit == 0 {
		limit = defaultLimit
	}

	repoNotes, err := s.noteRepo.GetList(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	var notes []*model.Note
	for _, repoNote := range repoNotes {
		serviceNote := converter.RepoNoteToServiceNote(repoNote)
		notes = append(notes, serviceNote)
	}

	return notes, nil
}
