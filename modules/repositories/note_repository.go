package repositories

import (
	"github.com/nattrio/go-clean-arch/modules/entities"
)

type NoteRepository interface {
	Save(note entities.Note) error
	Update(note entities.Note) error
	Delete(noteId int) error
	FindById(noteId int) (entities.Note, error)
	FindAll() ([]entities.Note, error)
}
