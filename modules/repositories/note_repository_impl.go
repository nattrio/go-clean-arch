package repositories

import (
	"github.com/nattrio/go-clean-arch/modules/entities"
	"github.com/nattrio/go-clean-arch/modules/models/request"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: db}
}

func (n *NoteRepositoryImpl) Save(note entities.Note) error {
	result := n.Db.Create(&note)
	return result.Error
}

func (n *NoteRepositoryImpl) Update(note entities.Note) error {
	var updateNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := n.Db.Model(&note).Updates(updateNote)
	return result.Error
}

func (n *NoteRepositoryImpl) Delete(noteId int) error {
	var note entities.Note
	result := n.Db.Where("id = ?", noteId).Delete(&note)
	return result.Error
}

func (n *NoteRepositoryImpl) FindById(noteId int) (entities.Note, error) {
	var note entities.Note
	result := n.Db.Where("id = ?", noteId).First(&note)
	return note, result.Error
}

func (n *NoteRepositoryImpl) FindAll() ([]entities.Note, error) {
	var notes []entities.Note
	result := n.Db.Find(&notes)
	return notes, result.Error
}
