package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sinameshkini/microkit/models"
	"lexilift/pkg/endpoint"
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

func (h *handler) startReview(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(endpoint.StartReviewRequest)
	)

	if err = c.BodyParser(req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.c.StartReview(ctx, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *handler) getReview(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		id  = models.Parse(c.Params("id"))
	)

	resp, meta, err := h.c.GetReview(ctx, id)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, meta)
}

func (h *handler) nextWord(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(endpoint.NextWordRequest)
		id  = models.Parse(c.Params("id"))
	)

	if err = c.BodyParser(req); err != nil {
		return responseError(c, err)
	}

	req.ReviewID = id

	resp, meta, err := h.c.NextWord(ctx, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, meta)
}
