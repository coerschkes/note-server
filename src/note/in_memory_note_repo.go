package note

import (
	"errors"
	"strconv"
)

type InMemoryNoteRepository struct {
	notes *[]Note
}

func MakeInMemoryNoteRepository() InMemoryNoteRepository {
	var notes = []Note{
		{ID: 1, Title: "Test", Content: "This is a test."},
		{ID: 2, Title: "Test2", Content: "This is the second test."},
	}
	return InMemoryNoteRepository{notes: &notes}
}

func (r InMemoryNoteRepository) GetAll() ([]Note, error) {
	return *r.notes, nil
}

func (r InMemoryNoteRepository) GetById(id int64) (Note, error) {
	for _, note := range *r.notes {
		if note.ID == id {
			return note, nil
		}
	}
	return Note{ID: -1}, errors.New("Note with id '" + strconv.Itoa(int(id)) + "' not found")
}

func (r InMemoryNoteRepository) Add(note Note) error {
	*r.notes = append(*r.notes, note)
	return nil
}

func (r InMemoryNoteRepository) DeleteById(id int64) error {
	for i, note := range *r.notes {
		if note.ID == id {
			*r.notes = append((*r.notes)[:i], (*r.notes)[i+1:]...)
			return nil
		}
	}
	return errors.New("Note with id '" + strconv.Itoa(int(id)) + "' not found")
}
