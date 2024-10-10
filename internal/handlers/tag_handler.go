package handlers

import (
	"github.com/burakaktna/VugoPress/pkg/utils"
	"strconv"

	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
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

	app.Get("/tags", tagHandler.GetTags)
	app.Post("/tags", jwtMiddleware, tagHandler.CreateTag)
	app.Get("/tags/:id", tagHandler.GetTag)
	app.Put("/tags/:id", jwtMiddleware, tagHandler.UpdateTag)
	app.Delete("/tags/:id", jwtMiddleware, tagHandler.DeleteTag)
}

func (h *TagHandler) GetTags(c *fiber.Ctx) error {
	tags, err := h.tagService.Index()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tags)
}

func (h *TagHandler) CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)
	if err := c.BodyParser(tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON ayrıştırılamıyor",
		})
	}

	var errors []utils.ErrorResponse
	errors = validator.Validate(tag)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	createdTag, err := h.tagService.Create(tag)
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

	tag, err := h.tagService.Show(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}
	return c.JSON(tag)
}

func (h *TagHandler) UpdateTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz etiket ID"})
	}

	updates := new(models.Tag)
	if err := c.BodyParser(updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON ayrıştırılamıyor",
		})
	}

	var errors []utils.ErrorResponse
	errors = validator.Validate(updates)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	updatedTag, err := h.tagService.Update(uint(id), updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Etiket bulunamadı"})
	}

	return c.JSON(updatedTag)
}

func (h *TagHandler) DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}
	err = h.tagService.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
