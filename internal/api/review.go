package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sinameshkini/microkit/models"
)

func (h *handler) fetchReviews(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(models.Request)
	)

	if err = c.QueryParser(req); err != nil {
		return responseError(c, err)
	}

	resp, meta, err := h.c.FetchReviews(ctx, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, meta)
}
