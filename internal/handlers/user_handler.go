package handlers

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"github.com/burakaktna/VugoPress/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func RegisterUserHandlers(app *fiber.App, userService services.UserService) {
	userHandler := NewUserHandler(userService)

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	userPost := new(models.UserPost)
	if err := c.BodyParser(userPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON ayrıştırılamıyor",
		})
	}

	var errors []utils.ErrorResponse
	errors = validator.Validate(userPost)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	registeredUser, err := h.userService.Register(userPost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(registeredUser)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	credentials := struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}{}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON ayrıştırılamıyor",
		})
	}

	var errors []utils.ErrorResponse
	errors = validator.Validate(credentials)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	token, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz e-posta veya şifre"})
	}

	return c.JSON(fiber.Map{"token": token})
}
