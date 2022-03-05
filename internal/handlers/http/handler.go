package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/breeders-zone/morphs-service/docs"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{

	}
}

func (h Handler) Init(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("pong")
	})
}