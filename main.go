package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-clean-arch/config"
	"github.com/nattrio/go-clean-arch/modules/deliveries"
	"github.com/nattrio/go-clean-arch/modules/deliveries/routes"
	"github.com/nattrio/go-clean-arch/modules/entities"
	"github.com/nattrio/go-clean-arch/modules/repositories"
	"github.com/nattrio/go-clean-arch/modules/services"
)

func main() {
	fmt.Println("Running server...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&entities.Note{})

	// Init Repository
	noteRepository := repositories.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := services.NewNoteServiceImpl(noteRepository, validate)

	// Init Controller
	noteHandler := deliveries.NewNoteHandler(noteService)

	// Routes
	routes := routes.NewRouter(noteHandler)

	// Run App
	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))
}
