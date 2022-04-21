package model

import (
	"encoding/json"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n Note) ToJson() (string, error) {
	if json, err := json.Marshal(n); err != nil {
		return "", err
	} else {
		return string(json), nil
	}
}

type NoteRepository interface {
	GetAll() ([]Note, error)
	GetById(id string) (Note, error)
	Add(note Note) error
	DeleteById(id string) error
}
