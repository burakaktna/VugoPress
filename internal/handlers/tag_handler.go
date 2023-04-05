package handlers

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TagHandler struct {
	tagService services.TagService
}

func NewTagHandler(tagService services.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func RegisterTagHandlers(app *fiber.App, tagService services.TagService, jwtMiddleware fiber.Handler) {
	tagHandler := NewTagHandler(tagService)

	app.Get("/tags", jwtMiddleware, tagHandler.GetTags)
	app.Post("/tags", jwtMiddleware, tagHandler.CreateTag)
	app.Get("/tags/:id", jwtMiddleware, tagHandler.GetTag)
	app.Put("/tags/:id", jwtMiddleware, tagHandler.UpdateTag)
	app.Delete("/tags/:id", jwtMiddleware, tagHandler.DeleteTag)
}

func (h *TagHandler) GetTags(c *fiber.Ctx) error {
	tags, err := h.tagService.GetTags()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tags)
}

func (h *TagHandler) CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)
	if err := c.BodyParser(tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdTag, err := h.tagService.CreateTag(tag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdTag)
}

func (h *TagHandler) GetTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	tag, err := h.tagService.GetTag(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}
	return c.JSON(tag)
}

func (h *TagHandler) UpdateTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	updates := new(models.Tag)
	if err := c.BodyParser(updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updatedTag, err := h.tagService.UpdateTag(uint(id), updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	return c.JSON(updatedTag)
}

func (h *TagHandler) DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}
	err = h.tagService.DeleteTag(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
