package converter

import (
	"github.com/Din4EE/note-service-api/internal/service/model"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToServiceNote(info *desc.NoteInfo) *model.Note {
	return &model.Note{
		NoteInfo: &model.NoteInfo{
			Title:  info.GetTitle(),
			Text:   info.GetText(),
			Author: info.GetAuthor(),
			Email:  info.GetEmail(),
		},
	}
}

func ToServiceNoteFromUpdate(info *desc.UpdateNoteInfo) *model.NoteInfoUpdate {
	note := &model.NoteInfoUpdate{}

	if info.GetTitle() != nil {
		note.Title = &info.GetTitle().Value
	}
	if info.GetText() != nil {
		note.Text = &info.GetText().Value
	}
	if info.GetAuthor() != nil {
		note.Author = &info.GetAuthor().Value
	}
	if info.GetEmail() != nil {
		note.Email = &info.GetEmail().Value
	}

	return note
}

func ToDescNote(serviceNote *model.Note) *desc.Note {
	note := &desc.Note{
		Id: serviceNote.ID,
		Info: &desc.NoteInfo{
			Title:  serviceNote.NoteInfo.Title,
			Text:   serviceNote.NoteInfo.Text,
			Author: serviceNote.NoteInfo.Author,
			Email:  serviceNote.NoteInfo.Email,
		},
		CreatedAt: timestamppb.New(serviceNote.CreatedAt),
	}

	if !serviceNote.UpdatedAt.IsZero() {
		note.UpdatedAt = timestamppb.New(*serviceNote.UpdatedAt)
	}

	return note
}
