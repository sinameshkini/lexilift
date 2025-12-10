package api

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) dashboard(c *fiber.Ctx) (err error) {
	resp, err := h.c.UserDashboard()
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}
