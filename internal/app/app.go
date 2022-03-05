package app

import (
	"fmt"
	"log"

	"github.com/breeders-zone/morphs-service/internal/config"
	"github.com/breeders-zone/morphs-service/internal/handlers/http"
	"github.com/gofiber/fiber/v2"
)

// @title Morphs service API
// @version 1.0
// @description Morphs service API

// @host localhost:3000
// @BasePath /

// Run initializes whole application.
func Run() {
	app := fiber.New()

	conf, err := config.Init(".")
	if err != nil {
		log.Fatalf("not load config")
	}

	fmt.Print(conf)

	h := http.NewHandler()

	h.Init(app)

	app.Listen(":3000")
}