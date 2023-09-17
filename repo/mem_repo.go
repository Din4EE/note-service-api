package repo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Note struct {
	ID     string
	Title  string
	Text   string
	Author string
}

type NoteRepository interface {
	Create(note Note) (string, error)
	Get(id string) (Note, error)
	GetList(limit int64, offset int64, query string) ([]Note, error)
	Update(id string, note Note) error
	Delete(id string) error
}

type InMemoryNoteRepository struct {
	data map[string]Note
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{data: make(map[string]Note)}
}

func (r *InMemoryNoteRepository) Create(note Note) (string, error) {
	id := uuid.New().String()
	note.ID = id
	r.data[id] = note
	return id, nil
}

func (r *InMemoryNoteRepository) Get(id string) (Note, error) {
	note, ok := r.data[id]
	if !ok {
		return Note{}, fmt.Errorf("note with id %s not found", id)
	}
	return note, nil
}

func (r *InMemoryNoteRepository) GetList(limit int64, offset int64, query string) ([]Note, error) {
	var notes []Note
	for _, note := range r.data {
		if strings.Contains(note.Title, query) || strings.Contains(note.Text, query) || strings.Contains(note.Author, query) {
			notes = append(notes, note)
		}
	}
	start := int(offset)
	end := int(limit + offset)
	if end > len(notes) {
		end = len(notes)
	}
	if start > end {
		start = end
	}

	return notes[start:end], nil
}

func (r *InMemoryNoteRepository) Update(id string, note Note) error {
	_, ok := r.data[id]
	if !ok {
		return fmt.Errorf("note with id %s not found", id)
	}
	note.ID = id
	r.data[id] = note
	return nil
}

func (r *InMemoryNoteRepository) Delete(id string) error {
	_, exists := r.data[id]
	if !exists {
		return fmt.Errorf("note with id %s not found", id)
	}
	delete(r.data, id)
	return nil
}

func (r *InMemoryNoteRepository) GetCurrentStatus() map[string]Note {
	return r.data
}
