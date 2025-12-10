package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sinameshkini/microkit/models"
	"lexilift/pkg/endpoint"
)

func (h *handler) fetchWords(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(models.Request)
	)

	if err = c.QueryParser(req); err != nil {
		return responseError(c, err)
	}

	resp, meta, err := h.c.FetchWords(ctx, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, meta)
}

func (h *handler) addWord(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(endpoint.AddWordRequest)
	)

	if err = c.BodyParser(req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.c.AddWord(ctx, req.Word)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}
