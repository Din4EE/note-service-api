package note_v1

import (
	"github.com/Din4EE/note-service-api/internal/service/note"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

type Note struct {
	noteService *note.Service
	desc.UnimplementedNoteServiceServer
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}
