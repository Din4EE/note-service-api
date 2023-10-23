package note

import (
	"context"
	"errors"
	"time"

	"github.com/Din4EE/note-service-api/internal/pkg/db"
	"github.com/Din4EE/note-service-api/internal/repo"
	"github.com/Din4EE/note-service-api/internal/repo/converter"
	repoModel "github.com/Din4EE/note-service-api/internal/repo/model"
	"github.com/Din4EE/note-service-api/internal/service/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
)

type repository struct {
	client db.Client
}

func NewRepository(db db.Client) repo.NoteRepository {
	return &repository{
		client: db,
	}
}

const tableName = "note"

func (r *repository) Create(ctx context.Context, note *model.Note) (uint64, error) {
	repoNote := converter.ToRepoNote(note)
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author", "email").
		Values(repoNote.NoteInfo.Title, repoNote.NoteInfo.Text, repoNote.NoteInfo.Author, repoNote.NoteInfo.Email).
		Suffix("returning id")

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.client.DB().QueryContext(ctx, db.Query{
		Name:     "CreateNote",
		QueryRaw: dbQuery,
	}, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, id uint64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "email", "created_at", "updated_at").
		From(tableName).
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	note := &repoModel.Note{}
	err = r.client.DB().GetContext(ctx, note, db.Query{
		Name:     "GetNote",
		QueryRaw: dbQuery,
	}, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToServiceNote(note), nil
}

func (r *repository) GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "email", "created_at", "updated_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Limit(limit).
		Offset(offset).
		OrderBy("id asc")
	if query != "" {
		builder = builder.Where(sq.Or{
			sq.Like{"title": "%" + query + "%"},
			sq.Like{"text": "%" + query + "%"},
			sq.Like{"author": "%" + query + "%"},
			sq.Like{"email": "%" + query + "%"},
		})
	}

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var notes []*repoModel.Note
	var serviceNotes []*model.Note
	err = r.client.DB().SelectContext(ctx, &notes, db.Query{
		Name:     "GetNoteList",
		QueryRaw: dbQuery,
	}, args...)
	if err != nil {
		return nil, err
	}

	for _, note := range notes {
		serviceNotes = append(serviceNotes, converter.ToServiceNote(note))
	}

	return serviceNotes, nil
}

func (r *repository) Update(ctx context.Context, note *model.Note) error {
	repoNote := converter.ToRepoNote(note)
	builder := sq.Update(tableName).Where(sq.Eq{"id": note.ID}).Set("updated_at", time.Now().UTC()).PlaceholderFormat(sq.Dollar)

	if repoNote.NoteInfo.Title.Valid {
		builder = builder.Set("title", repoNote.NoteInfo.Title)
	}
	if repoNote.NoteInfo.Text.Valid {
		builder = builder.Set("text", repoNote.NoteInfo.Text)
	}
	if repoNote.NoteInfo.Author.Valid {
		builder = builder.Set("author", repoNote.NoteInfo.Author)
	}
	if repoNote.NoteInfo.Email.Valid {
		builder = builder.Set("email", repoNote.NoteInfo.Email)
	}

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.client.DB().ExecContext(ctx, db.Query{
		Name:     "UpdateNote",
		QueryRaw: dbQuery,
	}, args...)
	if err != nil {
		return err
	}
	if count := res.RowsAffected(); count == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id uint64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id})

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.client.DB().ExecContext(ctx, db.Query{
		Name:     "DeleteNote",
		QueryRaw: dbQuery,
	}, args...)
	if err != nil {
		return err
	}
	if count := res.RowsAffected(); count == 0 {
		return errors.New("id not found")
	}

	return nil
}
