package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sinameshkini/microkit/models"
	"strings"

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

	app.Get("/audios/:name", h.getAudio)

	api := app.Group("/api/v1")

	api.Get("/status", h.status)

	api.Get("/dashboard", h.dashboard)

	// Words
	api.Get("/words", h.fetchWords)
	api.Post("/words", h.addWord)
	api.Get("/words/rand", h.addRandomWord)

	// Reviews
	api.Get("/reviews", h.fetchReviews)
	api.Post("/reviews", h.startReview)
	api.Get("/reviews/:id", h.getReview)
	api.Post("/reviews/:id/next", h.nextWord)

	return app.Listen(conf.ListenAddress)
}

func (h *handler) status(c *fiber.Ctx) (err error) {
	return response(c, map[string]string{"status": "running"}, nil)
}

func (h *handler) getAudio(c *fiber.Ctx) error {
	fileName := c.Params("name")
	filePath := fmt.Sprintf("./audios/%s", fileName)

	// Optional: prevent directory traversal attacks
	if strings.Contains(fileName, "..") {
		return c.Status(fiber.StatusBadRequest).SendString("invalid file name")
	}

	// Set correct Content-Type for MP3
	c.Type("mp3")

	return c.SendFile(filePath, true) // enable streaming
}

func response(c *fiber.Ctx, data, meta any) error {
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Data: data,
		Meta: meta,
	})
}

func responseError(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	if errors.Is(err, models.ErrNotfound) {
		status = fiber.StatusNotFound
	}
	return c.Status(status).JSON(models.Response{
		Message: err.Error(),
	})
}
