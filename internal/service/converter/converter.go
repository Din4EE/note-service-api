package converter

import (
	"github.com/Din4EE/note-service-api/internal/service/model"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToServiceNote(info *desc.NoteInfo) *model.Note {
	return &model.Note{
		NoteInfo: &model.NoteInfo{
			Title:  &info.Title,
			Text:   &info.Text,
			Author: &info.Author,
			Email:  &info.Email,
		},
	}
}

func ToServiceNoteFromUpdate(info *desc.UpdateNoteInfo) *model.Note {
	note := &model.Note{
		NoteInfo: &model.NoteInfo{},
	}

	if info.GetTitle() != nil {
		note.NoteInfo.Title = &info.Title.Value
	}
	if info.GetText() != nil {
		note.NoteInfo.Text = &info.Text.Value
	}
	if info.GetAuthor() != nil {
		note.NoteInfo.Author = &info.Author.Value
	}
	if info.GetEmail() != nil {
		note.NoteInfo.Email = &info.Email.Value
	}

	return note
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
