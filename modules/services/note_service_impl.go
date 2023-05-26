package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/nattrio/go-clean-arch/modules/entities"
	"github.com/nattrio/go-clean-arch/modules/models/request"
	"github.com/nattrio/go-clean-arch/modules/models/response"
	"github.com/nattrio/go-clean-arch/modules/repositories"
)

type NoteServiceImpl struct {
	NoteRepository repositories.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repositories.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		validate:       validate,
	}
}

func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) error {
	err := n.validate.Struct(note)
	if err != nil {
		return err
	}
	noteEntity := entities.Note{
		Content: note.Content,
	}
	n.NoteRepository.Save(noteEntity)
	return nil
}

func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) error {
	noteData, err := n.NoteRepository.FindById(note.Id)
	if err != nil {
		return err
	}
	noteData.Content = note.Content
	n.NoteRepository.Update(noteData)
	return nil
}

func (n *NoteServiceImpl) Delete(noteId int) error {
	err := n.NoteRepository.Delete(noteId)
	if err != nil {
		return err
	}
	return nil
}

func (n *NoteServiceImpl) FindById(noteId int) (*response.NoteResponse, error) {
	noteData, err := n.NoteRepository.FindById(noteId)
	if err != nil {
		return nil, err
	}
	noteResponse := response.NoteResponse{
		Id:      noteData.Id,
		Content: noteData.Content,
	}
	return &noteResponse, nil
}

func (n *NoteServiceImpl) FindAll() ([]response.NoteResponse, error) {
	result, err := n.NoteRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var notes []response.NoteResponse
	for _, note := range result {
		noteResponse := response.NoteResponse{
			Id:      note.Id,
			Content: note.Content,
		}
		notes = append(notes, noteResponse)
	}
	return notes, nil
}
