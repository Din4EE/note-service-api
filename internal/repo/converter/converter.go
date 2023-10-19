package converter

import (
	"database/sql"
	"time"

	repoModel "github.com/Din4EE/note-service-api/internal/repo/model"
	srvModel "github.com/Din4EE/note-service-api/internal/service/model"
)

func ServiceNoteToRepoNote(serviceNote *srvModel.Note) *repoModel.Note {
	return &repoModel.Note{
		ID:        serviceNote.ID,
		Title:     sql.NullString{String: *serviceNote.Title, Valid: *serviceNote.Title != ""},
		Text:      sql.NullString{String: *serviceNote.Text, Valid: *serviceNote.Text != ""},
		Author:    sql.NullString{String: *serviceNote.Author, Valid: *serviceNote.Author != ""},
		Email:     sql.NullString{String: *serviceNote.Email, Valid: *serviceNote.Email != ""},
		CreatedAt: serviceNote.CreatedAt,
		UpdatedAt: func(time *time.Time) sql.NullTime {
			if time == nil || time.IsZero() {
				return sql.NullTime{}
			}
			return sql.NullTime{Time: *time, Valid: true}
		}(serviceNote.UpdatedAt),
	}
}

func RepoNoteToServiceNote(repoNote *repoModel.Note) *srvModel.Note {
	return &srvModel.Note{
		ID:        repoNote.ID,
		Title:     &repoNote.Title.String,
		Text:      &repoNote.Text.String,
		Author:    &repoNote.Author.String,
		Email:     &repoNote.Email.String,
		CreatedAt: repoNote.CreatedAt,
		UpdatedAt: func(nt sql.NullTime) *time.Time {
			if !nt.Valid {
				return &time.Time{}
			}
			return &nt.Time
		}(repoNote.UpdatedAt),
	}
}
