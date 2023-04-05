package handlers

import (
	"github.com/burakaktna/VugoPress/internal/config"
	"github.com/burakaktna/VugoPress/internal/repository"
	"github.com/burakaktna/VugoPress/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func RegisterHandlers(app *fiber.App, db *gorm.DB, jwtMiddleware fiber.Handler) {
	cfg := config.New()

	articleRepo := repository.NewArticleRepository(db)
	articleService := services.NewArticleService(articleRepo)
	RegisterArticleHandlers(app, articleService, jwtMiddleware)

	emailService := services.NewEmailService(cfg)
	contactRepo := repository.NewContactRepository(db)
	contactService := services.NewContactService(contactRepo, emailService)
	RegisterContactHandlers(app, contactService, jwtMiddleware)

	usefulLinkRepo := repository.NewUsefulLinkRepository(db)
	usefulLinkService := services.NewUsefulLinksService(usefulLinkRepo)
	RegisterUsefulLinkHandlers(app, usefulLinkService, jwtMiddleware)

	tagRepo := repository.NewTagRepository(db)
	tagService := services.NewTagService(tagRepo)
	RegisterTagHandlers(app, tagService, jwtMiddleware)

	socialMediaLinkRepo := repository.NewSocialMediaLinkRepository(db)
	socialMediaLinkService := services.NewSocialMediaLinkService(socialMediaLinkRepo)
	RegisterSocialMediaLinkHandlers(app, socialMediaLinkService, jwtMiddleware)
}
