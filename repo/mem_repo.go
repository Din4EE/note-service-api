package repo

import (
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
)

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

type InMemoryNoteRepository struct {
	data map[string]Note
	mu   sync.RWMutex
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{data: make(map[string]Note), mu: sync.RWMutex{}}
}

func (r *InMemoryNoteRepository) Create(note Note) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()
	note.ID = id
	r.data[id] = note

	return id, nil
}

func (r *InMemoryNoteRepository) Get(id string) (Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, ok := r.data[id]
	if !ok {
		return Note{}, fmt.Errorf("note with id %s not found", id)
	}

	return note, nil
}

func (r *InMemoryNoteRepository) GetList(limit int64, offset int64, query string) ([]Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

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

func (r *InMemoryNoteRepository) Update(id string, note NoteUpdate) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	noteData, ok := r.data[id]
	if !ok {
		return fmt.Errorf("note with id %s not found", id)
	}

	if note.Title != nil {
		noteData.Title = *note.Title
	}
	if note.Text != nil {
		noteData.Text = *note.Text
	}
	if note.Author != nil {
		noteData.Author = *note.Author
	}
	r.data[id] = noteData

	return nil
}

func (r *InMemoryNoteRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[id]
	if !ok {
		return fmt.Errorf("note with id %s not found", id)
	}
	delete(r.data, id)

	return nil
}
