package note

import (
	"github.com/Din4EE/note-service-api/internal/repo"
)

type Service struct {
	noteRepo repo.NoteRepository
}

func NewService(noteRepository repo.NoteRepository) *Service {
	return &Service{
		noteRepo: noteRepository,
	}
}
