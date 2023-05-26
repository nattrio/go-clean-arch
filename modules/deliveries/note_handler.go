package deliveries

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-clean-arch/modules/models/request"
	"github.com/nattrio/go-clean-arch/modules/models/response"
	"github.com/nattrio/go-clean-arch/modules/services"
)

type NoteHandler struct {
	noteService services.NoteService
}

func NewNoteHandler(noteService services.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
	}
}

func (h *NoteHandler) CreateNote(ctx *fiber.Ctx) error {
	createNoteRequest := request.CreateNoteRequest{}
	if err := ctx.BodyParser(&createNoteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	h.noteService.Create(createNoteRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully create note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (h *NoteHandler) UpdateNote(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	if err := ctx.BodyParser(&updateNoteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	updateNoteRequest.Id = id
	h.noteService.Update(updateNoteRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully update note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (h *NoteHandler) DeleteNote(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	h.noteService.Delete(id)
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully delete note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (h *NoteHandler) FindNoteById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	note, err := h.noteService.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully find note",
		Data:    note,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (h *NoteHandler) FindAllNotes(ctx *fiber.Ctx) error {
	notes, err := h.noteService.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "ok",
		Message: "successfully find all notes",
		Data:    notes,
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
