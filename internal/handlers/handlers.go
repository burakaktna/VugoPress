package handlers

import (
	"github.com/burakaktna/VugoPress/internal/config"
	"github.com/burakaktna/VugoPress/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func RegisterHandlers(app *fiber.App, db *gorm.DB, jwtMiddleware fiber.Handler) {
	cfg := config.New()

	userService := services.NewUserService(db, cfg.AppKey)
	RegisterUserHandlers(app, userService)

	articleService := services.NewArticleService(db)
	RegisterArticleHandlers(app, articleService, jwtMiddleware)

	emailService := services.NewEmailService(cfg)
	contactService := services.NewContactService(db, emailService)
	RegisterContactHandlers(app, contactService, jwtMiddleware)

	usefulLinkService := services.NewUsefulLinkService(db)
	RegisterUsefulLinkHandlers(app, usefulLinkService, jwtMiddleware)

	tagService := services.NewTagService(db)
	RegisterTagHandlers(app, tagService, jwtMiddleware)

	socialMediaLinkService := services.NewSocialMediaLinkService(db)
	RegisterSocialMediaLinkHandlers(app, socialMediaLinkService, jwtMiddleware)
}
