package ram

import (
	"co/note-server/src/domain/model"
	"errors"
)

type InMemoryNoteRepository struct {
	notes *[]model.Note
}

func MakeInMemoryNoteRepository() InMemoryNoteRepository {
	var notes = []model.Note{
		{ID: "1", Title: "Test", Content: "This is a test."},
		{ID: "2", Title: "Test2", Content: "This is the second test."},
	}
	return InMemoryNoteRepository{notes: &notes}
}

func (r InMemoryNoteRepository) GetAll() ([]model.Note, error) {
	return *r.notes, nil
}

func (r InMemoryNoteRepository) GetById(id string) (model.Note, error) {
	for _, note := range *r.notes {
		if note.ID == id {
			return note, nil
		}
	}
	return model.MakeInvalidNote(), errors.New("Note with id '" + id + "' not found")
}

func (r InMemoryNoteRepository) Add(note model.Note) error {
	if n, _ := r.GetById(note.ID); n.ID == model.MakeInvalidNote().ID {
		*r.notes = append(*r.notes, note)
		return nil
	} else {
		return errors.New("Note with id '" + note.ID + "' already exists.")
	}
}

func (r InMemoryNoteRepository) DeleteById(id string) error {
	for i, note := range *r.notes {
		if note.ID == id {
			*r.notes = append((*r.notes)[:i], (*r.notes)[i+1:]...)
			return nil
		}
	}
	return errors.New("Note with id '" + id + "' not found")
}
