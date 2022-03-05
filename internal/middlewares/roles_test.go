package middlewares_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/breeders-zone/morphs-service/internal/handlers/http/errors"
	"github.com/breeders-zone/morphs-service/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_RoleMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(middlewares.Role([]string{"user", "admin"})).Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	req := httptest.NewRequest("GET", "/ping", nil)
	req.Header.Add("X-User-Roles", "user,admin")
	

	res, _ := app.Test(req, -1)

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equalf(t, res.StatusCode, fiber.StatusOK, "TEST STATUS")
	assert.Equalf(t, string(body), "pong", "TEST BODY")
}

func Test_RoleMiddleware_Err(t *testing.T) {
	app := fiber.New()

	app.Use(middlewares.Role([]string{"user", "admin"})).Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	req := httptest.NewRequest("GET", "/ping", nil)
	

	res, _ := app.Test(req, -1)

	body, _ := ioutil.ReadAll(res.Body)

	var errResp errors.ErrorResponse
	json.Unmarshal(body, &errResp)



	assert.Equalf(t, res.StatusCode, fiber.StatusForbidden, "TEST STATUS")
	assert.Equalf(t, errResp, errors.ErrorResponse{fiber.StatusForbidden, "Forbidden"}, "TEST BODY")
}