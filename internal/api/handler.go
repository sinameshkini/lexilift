package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sinameshkini/microkit/models"

	"lexilift/internal/config"
	"lexilift/internal/core"
)

type handler struct {
	c *core.Core
}

func Init(conf *config.Config, c *core.Core) (err error) {
	h := &handler{c: c}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/v1")

	api.Get("/status", h.status)

	return app.Listen(conf.ListenAddress)
}

func (h *handler) status(c *fiber.Ctx) (err error) {
	return response(c, map[string]string{"status": "running"}, nil)
}

func response(c *fiber.Ctx, data, meta any) error {
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Data: data,
		Meta: meta,
	})
}

func responseError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
		Message: err.Error(),
	})
}
