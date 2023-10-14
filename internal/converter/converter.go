package converter

import (
	"database/sql"
	"time"

	"github.com/Din4EE/note-service-api/internal/repo"
	"github.com/Din4EE/note-service-api/internal/service"
	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ServiceNoteToRepoNote(serviceNote *service.Note) *repo.Note {
	return &repo.Note{
		ID:        serviceNote.ID,
		Title:     sql.NullString{String: serviceNote.Title, Valid: serviceNote.Title != ""},
		Text:      sql.NullString{String: serviceNote.Text, Valid: serviceNote.Text != ""},
		Author:    sql.NullString{String: serviceNote.Author, Valid: serviceNote.Author != ""},
		Email:     sql.NullString{String: serviceNote.Email, Valid: serviceNote.Email != ""},
		CreatedAt: serviceNote.CreatedAt,
		UpdatedAt: func(time time.Time) sql.NullTime {
			if time.IsZero() {
				return sql.NullTime{}
			}
			return sql.NullTime{Time: time, Valid: true}
		}(serviceNote.UpdatedAt),
	}
}

func RepoNoteToServiceNote(repoNote *repo.Note) *service.Note {
	return &service.Note{
		ID:        repoNote.ID,
		Title:     repoNote.Title.String,
		Text:      repoNote.Text.String,
		Author:    repoNote.Author.String,
		Email:     repoNote.Email.String,
		CreatedAt: repoNote.CreatedAt,
		UpdatedAt: func(nt sql.NullTime) time.Time {
			if !nt.Valid {
				return time.Time{}
			}
			return nt.Time
		}(repoNote.UpdatedAt),
	}
}

func UpdateRequestToRepoNote(req *desc.UpdateNoteRequest) *repo.Note {
	return &repo.Note{
		Title:  sql.NullString{String: req.GetTitle().GetValue(), Valid: req.GetTitle() != nil},
		Text:   sql.NullString{String: req.GetText().GetValue(), Valid: req.GetText() != nil},
		Author: sql.NullString{String: req.GetAuthor().GetValue(), Valid: req.GetAuthor() != nil},
		Email:  sql.NullString{String: req.GetEmail().GetValue(), Valid: req.GetEmail() != nil},
	}
}

func ServiceNoteToDescNote(serviceNote *service.Note) *desc.Note {
	note := &desc.Note{
		Id:        serviceNote.ID,
		Title:     serviceNote.Title,
		Text:      serviceNote.Text,
		Author:    serviceNote.Author,
		Email:     serviceNote.Email,
		CreatedAt: timestamppb.New(serviceNote.CreatedAt),
	}

	if !serviceNote.UpdatedAt.IsZero() {
		note.UpdatedAt = timestamppb.New(serviceNote.UpdatedAt)
	}
	return note
}
