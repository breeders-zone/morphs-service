package http_test

import (
	"net/http/httptest"
	"testing"

	handler "github.com/breeders-zone/morphs-service/internal/handlers/http"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewHandler(t *testing.T) {
	h := handler.NewHandler()

	require.IsType(t, &handler.Handler{}, h)
}

func Test_NewHandler_Routing(t *testing.T) {
	app := fiber.New()

	h := handler.NewHandler()
	h.Init(app)

	req := httptest.NewRequest("GET", "/ping", nil)
	

	res, _ := app.Test(req, -1)

	assert.Equalf(t, res.StatusCode, fiber.StatusOK, "Test ping")
}