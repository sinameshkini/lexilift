package api

import (
	"github.com/gofiber/fiber/v2"
	"lexilift/pkg/endpoint"
)

func (h *handler) writingChallenge(c *fiber.Ctx) (err error) {
	var (
		ctx = c.Context()
		req = new(endpoint.ChallengeRequest)
	)

	if err = c.BodyParser(req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.c.Challenge(ctx, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}
