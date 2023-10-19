package note

import (
	"context"
	"time"

	"github.com/Din4EE/note-service-api/internal/config"
	"github.com/Din4EE/note-service-api/internal/repo"
	"github.com/Din4EE/note-service-api/internal/repo/converter"
	repoModel "github.com/Din4EE/note-service-api/internal/repo/model"
	"github.com/Din4EE/note-service-api/internal/service/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(cfg config.PgConfig) repo.NoteRepository {
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		panic(err)
	}

	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, note *model.Note) (uint64, error) {
	repoNote := converter.ServiceNoteToRepoNote(note)
	builder := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author", "email").
		Values(repoNote.Title, repoNote.Text, repoNote.Author, repoNote.Email).
		Suffix("returning id")

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, dbQuery, args...)
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
		From("note").
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowxContext(ctx, dbQuery, args...)
	note := &repoModel.Note{}
	err = row.Scan(&note.ID, &note.Title, &note.Text, &note.Author, &note.Email, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.RepoNoteToServiceNote(note), nil
}

func (r *repository) GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "email", "created_at", "updated_at").
		From("note").
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

	rows, err := r.db.QueryxContext(ctx, dbQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*model.Note
	for rows.Next() {
		note := &repoModel.Note{}
		err = rows.Scan(&note.ID, &note.Title, &note.Text, &note.Author, &note.Email, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, converter.RepoNoteToServiceNote(note))
	}

	return notes, nil
}

func (r *repository) Update(ctx context.Context, note *model.Note) error {
	repoNote := converter.ServiceNoteToRepoNote(note)
	builder := sq.Update("note").Where(sq.Eq{"id": note.ID}).Set("updated_at", time.Now().UTC()).PlaceholderFormat(sq.Dollar)

	if repoNote.Title.Valid {
		builder = builder.Set("title", repoNote.Title)
	}
	if repoNote.Text.Valid {
		builder = builder.Set("text", repoNote.Text)
	}
	if repoNote.Author.Valid {
		builder = builder.Set("author", repoNote.Author)
	}
	if repoNote.Email.Valid {
		builder = builder.Set("email", repoNote.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id uint64) error {
	builder := sq.Delete("note").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id})

	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, dbQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
