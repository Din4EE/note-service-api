package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Din4EE/note-service-api/config"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresNoteRepository struct {
	db *sqlx.DB
}

func NewPostgresNoteRepository(cfg config.PgConfig) *PostgresNoteRepository {
	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Connect("pgx", dbDsn)
	if err != nil {
		panic(err)
	}

	return &PostgresNoteRepository{
		db: db,
	}
}

func (r *PostgresNoteRepository) Create(note Note) (string, error) {
	builder := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("title", "text", "author").
		Values(note.Title, note.Text, note.Author).
		Suffix("returning id")
	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}
	row, err := r.db.QueryContext(context.Background(), dbQuery, args...)
	if err != nil {
		return "", err
	}
	defer row.Close()

	row.Next()
	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *PostgresNoteRepository) Get(id string) (Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		From("note").
		Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)
	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return Note{}, err
	}
	row := r.db.QueryRowxContext(context.Background(), dbQuery, args...)

	var note Note
	err = row.Scan(&note.ID, &note.Title, &note.Text, &note.Author)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (r *PostgresNoteRepository) GetList(limit int64, offset int64, query string) ([]Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		From("note").
		Where(sq.Or{
			sq.Like{"title": query},
			sq.Like{"text": query},
			sq.Like{"author": query},
		}).
		PlaceholderFormat(sq.Dollar).
		Limit(uint64(limit)).
		Offset(uint64(offset))
	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryxContext(context.Background(), dbQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.ID, &note.Title, &note.Text, &note.Author)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *PostgresNoteRepository) Update(id string, note NoteUpdate) error {
	builder := sq.Update("note").Where(sq.Eq{"id": id}).Set("updated_at", time.Now()).PlaceholderFormat(sq.Dollar)

	if note.Title != nil {
		builder = builder.Set("title", *note.Title)
	}
	if note.Text != nil {
		builder = builder.Set("text", *note.Text)
	}
	if note.Author != nil {
		builder = builder.Set("author", *note.Author)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresNoteRepository) Delete(id string) error {
	builder := sq.Delete("note").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id})
	dbQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(context.Background(), dbQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
