package services

import (
	"github.com/nattrio/go-clean-arch/modules/models/request"
	"github.com/nattrio/go-clean-arch/modules/models/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest) error
	Update(note request.UpdateNoteRequest) error
	Delete(noteId int) error
	FindById(noteId int) (*response.NoteResponse, error)
	FindAll() ([]response.NoteResponse, error)
}
