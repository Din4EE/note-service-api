package converter

import (
	"github.com/Din4EE/note-service-api/internal/service/model"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToServiceNote(descNote *desc.Note) *model.Note {
	return &model.Note{
		ID: descNote.GetId(),
		NoteInfo: &model.NoteInfo{
			Title:  &descNote.GetInfo().Title,
			Text:   &descNote.GetInfo().Text,
			Author: &descNote.GetInfo().Author,
			Email:  &descNote.GetInfo().Email,
		},
	}
}

func ToDescNote(serviceNote *model.Note) *desc.Note {
	note := &desc.Note{
		Id: serviceNote.ID,
		Info: &desc.NoteInfo{
			Title:  *serviceNote.NoteInfo.Title,
			Text:   *serviceNote.NoteInfo.Text,
			Author: *serviceNote.NoteInfo.Author,
			Email:  *serviceNote.NoteInfo.Email,
		},
		CreatedAt: timestamppb.New(serviceNote.CreatedAt),
	}

	if !serviceNote.UpdatedAt.IsZero() {
		note.UpdatedAt = timestamppb.New(*serviceNote.UpdatedAt)
	}

	return note
}
