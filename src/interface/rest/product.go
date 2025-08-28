package rest

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ProductRoutes(router *fiber.App) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"data": map[string]any{
				"name": "arash",
			},
			"message": true,
		})
	})
}
