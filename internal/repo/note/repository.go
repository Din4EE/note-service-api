package note

import (
	"context"
	"time"

	"github.com/Din4EE/note-service-api/internal/config"
	"github.com/Din4EE/note-service-api/internal/repo"
	"github.com/Din4EE/note-service-api/internal/repo/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgresNoteRepository struct {
	db *sqlx.DB
}

func NewPostgresNoteRepository(cfg config.PgConfig) repo.NoteRepository {
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		panic(err)
	}

	return &postgresNoteRepository{
		db: db,
	}
}

func (r *postgresNoteRepository) Create(ctx context.Context, note *model.Note) (uint64, error) {
	builder := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author", "email").
		Values(note.Title, note.Text, note.Author, note.Email).
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

func (r *postgresNoteRepository) Get(ctx context.Context, id uint64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "email", "created_at", "updated_at").
		From("note").
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)
	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRowxContext(ctx, dbQuery, args...)
	note := &model.Note{}
	err = row.Scan(&note.ID, &note.Title, &note.Text, &note.Author, &note.Email, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *postgresNoteRepository) GetList(ctx context.Context, query string, limit uint64, offset uint64) ([]*model.Note, error) {
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
		note := &model.Note{}
		err = rows.Scan(&note.ID, &note.Title, &note.Text, &note.Author, &note.Email, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *postgresNoteRepository) Update(ctx context.Context, note *model.Note) error {
	builder := sq.Update("note").Where(sq.Eq{"id": note.ID}).Set("updated_at", time.Now().UTC()).PlaceholderFormat(sq.Dollar)

	if note.Title.Valid {
		builder = builder.Set("title", note.Title)
	}
	if note.Text.Valid {
		builder = builder.Set("text", note.Text)
	}
	if note.Author.Valid {
		builder = builder.Set("author", note.Author)
	}
	if note.Email.Valid {
		builder = builder.Set("email", note.Email)
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

func (r *postgresNoteRepository) Delete(ctx context.Context, id uint64) error {
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
