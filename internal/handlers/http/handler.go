package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/breeders-zone/morphs-service/docs"
	"github.com/breeders-zone/morphs-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h Handler) Init(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("pong")
	})

	app.Get("/genes/:id", h.GetGene)
	app.Post("/genes", h.CreateGene)
	app.Put("/genes/:id", h.UpdateGene)
	app.Delete("/genes/:id", h.DeleteGene)
}
