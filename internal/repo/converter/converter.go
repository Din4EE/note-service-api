package converter

import (
	"database/sql"
	"time"

	repoModel "github.com/Din4EE/note-service-api/internal/repo/model"
	srvModel "github.com/Din4EE/note-service-api/internal/service/model"
)

func ToRepoNote(serviceNote *srvModel.Note) *repoModel.Note {
	repoNote := &repoModel.Note{
		ID:        serviceNote.ID,
		CreatedAt: serviceNote.CreatedAt,
		NoteInfo: &repoModel.NoteInfo{
			Title:  sql.NullString{String: serviceNote.NoteInfo.Title, Valid: serviceNote.NoteInfo.Title != ""},
			Text:   sql.NullString{String: serviceNote.NoteInfo.Text, Valid: serviceNote.NoteInfo.Text != ""},
			Author: sql.NullString{String: serviceNote.NoteInfo.Author, Valid: serviceNote.NoteInfo.Author != ""},
			Email:  sql.NullString{String: serviceNote.NoteInfo.Email, Valid: serviceNote.NoteInfo.Email != ""},
		},
	}
	repoNote.UpdatedAt = func(time *time.Time) sql.NullTime {
		if time == nil || time.IsZero() {
			return sql.NullTime{}
		}
		return sql.NullTime{Time: *time, Valid: true}
	}(serviceNote.UpdatedAt)

	return repoNote
}

func ToRepoNoteFromUpdate(serviceNote *srvModel.NoteInfoUpdate) *repoModel.Note {
	repoNote := &repoModel.Note{
		NoteInfo: &repoModel.NoteInfo{},
	}

	if serviceNote.Title != nil {
		repoNote.NoteInfo.Title = sql.NullString{String: *serviceNote.Title, Valid: true}
	}
	if serviceNote.Text != nil {
		repoNote.NoteInfo.Text = sql.NullString{String: *serviceNote.Text, Valid: true}
	}
	if serviceNote.Author != nil {
		repoNote.NoteInfo.Author = sql.NullString{String: *serviceNote.Author, Valid: true}
	}
	if serviceNote.Email != nil {
		repoNote.NoteInfo.Email = sql.NullString{String: *serviceNote.Email, Valid: true}
	}

	return repoNote
}

func ToServiceNote(repoNote *repoModel.Note) *srvModel.Note {
	return &srvModel.Note{
		ID: repoNote.ID,
		NoteInfo: &srvModel.NoteInfo{
			Title:  repoNote.NoteInfo.Title.String,
			Text:   repoNote.NoteInfo.Text.String,
			Author: repoNote.NoteInfo.Author.String,
			Email:  repoNote.NoteInfo.Email.String,
		},
		CreatedAt: repoNote.CreatedAt,
		UpdatedAt: func(nt sql.NullTime) *time.Time {
			if !nt.Valid {
				return &time.Time{}
			}
			return &nt.Time
		}(repoNote.UpdatedAt),
	}
}
