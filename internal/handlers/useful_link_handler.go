package handlers

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UsefulLinkHandler struct {
	usefulLinkService services.UsefulLinksService
}

func NewUsefulLinkHandler(usefulLinkService services.UsefulLinksService) *UsefulLinkHandler {
	return &UsefulLinkHandler{usefulLinkService: usefulLinkService}
}

func RegisterUsefulLinkHandlers(app *fiber.App, usefulLinkService services.UsefulLinksService, jwtMiddleware fiber.Handler) {
	usefulLinkHandler := NewUsefulLinkHandler(usefulLinkService)

	app.Get("/useful_links", jwtMiddleware, usefulLinkHandler.GetUsefulLinks)
	app.Post("/useful_links", jwtMiddleware, usefulLinkHandler.CreateUsefulLink)
	app.Get("/useful_links/:id", jwtMiddleware, usefulLinkHandler.GetUsefulLink)
	app.Put("/useful_links/:id", jwtMiddleware, usefulLinkHandler.UpdateUsefulLink)
	app.Delete("/useful_links/:id", jwtMiddleware, usefulLinkHandler.DeleteUsefulLink)
}

func (h *UsefulLinkHandler) GetUsefulLinks(c *fiber.Ctx) error {
	usefulLinks, err := h.usefulLinkService.GetUsefulLinks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(usefulLinks)
}

func (h *UsefulLinkHandler) CreateUsefulLink(c *fiber.Ctx) error {
	usefulLink := new(models.UsefulLink)
	if err := c.BodyParser(usefulLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdUsefulLink, err := h.usefulLinkService.CreateUsefulLink(usefulLink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdUsefulLink)
}

func (h *UsefulLinkHandler) GetUsefulLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid useful link ID"})
	}

	usefulLink, err := h.usefulLinkService.GetUsefulLink(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Useful link not found"})
	}
	return c.JSON(usefulLink)
}

func (h *UsefulLinkHandler) UpdateUsefulLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid useful link ID"})
	}

	updates := new(models.UsefulLink)
	if err := c.BodyParser(updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updatedUsefulLink, err := h.usefulLinkService.UpdateUsefulLink(uint(id), updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Useful link not found"})
	}

	return c.JSON(updatedUsefulLink)
}

func (h *UsefulLinkHandler) DeleteUsefulLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid useful link ID"})
	}
	err = h.usefulLinkService.DeleteUsefulLink(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Useful link not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
