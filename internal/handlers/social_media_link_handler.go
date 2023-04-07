package handlers

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SocialMediaLinkHandler struct {
	socialMediaLinkService services.SocialMediaLinkService
}

func NewSocialMediaLinkHandler(service services.SocialMediaLinkService) *SocialMediaLinkHandler {
	return &SocialMediaLinkHandler{socialMediaLinkService: service}
}

func RegisterSocialMediaLinkHandlers(app *fiber.App, socialMediaLinkService services.SocialMediaLinkService, jwtMiddleware fiber.Handler) {
	socialMediaLinkHandler := NewSocialMediaLinkHandler(socialMediaLinkService)

	app.Get("/social_media_links", socialMediaLinkHandler.GetSocialMediaLinks)
	app.Post("/social_media_links", jwtMiddleware, socialMediaLinkHandler.CreateSocialMediaLink)
	app.Get("/social_media_links/:id", socialMediaLinkHandler.GetSocialMediaLink)
	app.Put("/social_media_links/:id", jwtMiddleware, socialMediaLinkHandler.UpdateSocialMediaLink)
	app.Delete("/social_media_links/:id", jwtMiddleware, socialMediaLinkHandler.DeleteSocialMediaLink)
}

func (h *SocialMediaLinkHandler) GetSocialMediaLinks(c *fiber.Ctx) error {
	socialMediaLinks, err := h.socialMediaLinkService.GetSocialMediaLinks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(socialMediaLinks)
}

func (h *SocialMediaLinkHandler) CreateSocialMediaLink(c *fiber.Ctx) error {
	socialMediaLink := new(models.SocialMediaLink)
	if err := c.BodyParser(socialMediaLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdSocialMediaLink, err := h.socialMediaLinkService.CreateSocialMediaLink(socialMediaLink)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdSocialMediaLink)
}

func (h *SocialMediaLinkHandler) GetSocialMediaLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid social media link ID"})
	}

	socialMediaLink, err := h.socialMediaLinkService.GetSocialMediaLink(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Social media link not found"})
	}
	return c.JSON(socialMediaLink)
}

func (h *SocialMediaLinkHandler) UpdateSocialMediaLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid social media link ID"})
	}

	updates := new(models.SocialMediaLink)
	if err := c.BodyParser(updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updatedSocialMediaLink, err := h.socialMediaLinkService.UpdateSocialMediaLink(uint(id), updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Social media link not found"})
	}

	return c.JSON(updatedSocialMediaLink)
}

func (h *SocialMediaLinkHandler) DeleteSocialMediaLink(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid social media link ID"})
	}
	err = h.socialMediaLinkService.DeleteSocialMediaLink(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Social media link not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
