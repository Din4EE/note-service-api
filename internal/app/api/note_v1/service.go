package note_v1

import (
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"github.com/Din4EE/note-service-api/repo"
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
