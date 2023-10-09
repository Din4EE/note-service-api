package note_v1

import (
	"github.com/Din4EE/note-service-api/internal/app/repo"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

type Note struct {
	repo repo.NoteRepository
	desc.UnimplementedNoteServiceServer
}

func NewNote(repo repo.NoteRepository) *Note {
	return &Note{
		repo: repo,
	}
}
