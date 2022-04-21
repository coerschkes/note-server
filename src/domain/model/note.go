package model

type Note struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteRepository interface {
	GetAll() ([]Note, error)
	GetById(id int64) (Note, error)
	Add(note Note) error
	DeleteById(id int64) error
}
