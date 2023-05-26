package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-clean-arch/modules/deliveries"
)

func NewRouter(noteHandler *deliveries.NoteHandler) *fiber.App {
	router := fiber.New()

	router.Get("/heathchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to server",
		})
	})

	router.Route("/notes", func(route fiber.Router) {
		route.Post("/", noteHandler.CreateNote)
		route.Get("/", noteHandler.FindAllNotes)
	})

	router.Route("/notes/:noteId", func(route fiber.Router) {
		route.Get("/", noteHandler.FindNoteById)
		route.Patch("/", noteHandler.UpdateNote)
		route.Delete("/", noteHandler.DeleteNote)
	})

	return router
}
