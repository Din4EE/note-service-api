package converter

import (
	"github.com/Din4EE/note-service-api/internal/service/model"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DescNoteToServiceNote(descNote *desc.Note) *model.Note {
	return &model.Note{
		ID:     descNote.Id,
		Title:  &descNote.Title,
		Text:   &descNote.Text,
		Author: &descNote.Author,
		Email:  &descNote.Email,
	}
}

func ServiceNoteToDescNote(serviceNote *model.Note) *desc.Note {
	note := &desc.Note{
		Id:        serviceNote.ID,
		Title:     *serviceNote.Title,
		Text:      *serviceNote.Text,
		Author:    *serviceNote.Author,
		Email:     *serviceNote.Email,
		CreatedAt: timestamppb.New(serviceNote.CreatedAt),
	}

	if !serviceNote.UpdatedAt.IsZero() {
		note.UpdatedAt = timestamppb.New(*serviceNote.UpdatedAt)
	}

	return note
}
