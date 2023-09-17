package note_v1

import (
	"context"

	desc "github.com/Din4EE/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	notes, err := n.repo.GetList(req.GetLimit(), req.GetOffset(), req.GetSearchQuery())
	if err != nil {
		return nil, err
	}
	var ids, titles, texts, authors []string
	for _, note := range notes {
		ids = append(ids, note.ID)
		titles = append(titles, note.Title)
		texts = append(texts, note.Text)
		authors = append(authors, note.Author)
	}
	return &desc.GetListNoteResponse{
		Ids:     ids,
		Titles:  titles,
		Texts:   texts,
		Authors: authors,
	}, nil
}
