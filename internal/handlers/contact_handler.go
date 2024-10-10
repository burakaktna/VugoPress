package handlers

import (
	"strconv"

	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"github.com/burakaktna/VugoPress/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var validator = &utils.XValidator{}

type ContactHandler struct {
	contactService services.ContactService
}

func NewContactHandler(contactService services.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func RegisterContactHandlers(app *fiber.App, contactService services.ContactService, jwtMiddleware fiber.Handler) {
	contactHandler := NewContactHandler(contactService)

	app.Get("/contacts", jwtMiddleware, contactHandler.GetContacts)
	app.Post("/contacts", contactHandler.CreateContact)
	app.Get("/contacts/:id", jwtMiddleware, contactHandler.GetContact)
	app.Put("/contacts/:id", jwtMiddleware, contactHandler.UpdateContact)
	app.Delete("/contacts/:id", jwtMiddleware, contactHandler.DeleteContact)
}

func (h *ContactHandler) GetContacts(c *fiber.Ctx) error {
	contacts, err := h.contactService.Index()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(contacts)
}

func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var contact models.Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var errors []utils.ErrorResponse
	errors = validator.Validate(contact)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	createdContact, err := h.contactService.Create(&contact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdContact)
}

func (h *ContactHandler) GetContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contact ID"})
	}

	contact, err := h.contactService.Show(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contact not found"})
	}
	return c.JSON(contact)
}

func (h *ContactHandler) UpdateContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contact ID"})
	}

	updates, err := h.contactService.Show(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}
	updatedContact, err := h.contactService.Update(updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Update is unsuccessful"})
	}

	return c.JSON(updatedContact)
}

func (h *ContactHandler) DeleteContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contact ID"})
	}
	err = h.contactService.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contact not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
