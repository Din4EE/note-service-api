package repo

type Note struct {
	ID     string
	Title  string
	Text   string
	Author string
}

type NoteUpdate struct {
	Title  *string
	Text   *string
	Author *string
}

type NoteRepository interface {
	Create(note Note) (string, error)
	Get(id string) (Note, error)
	GetList(limit int64, offset int64, query string) ([]Note, error)
	Update(id string, note NoteUpdate) error
	Delete(id string) error
}
