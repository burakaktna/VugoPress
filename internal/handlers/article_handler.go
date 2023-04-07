// internal/handlers/article_handler.go
package handlers

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	articleService services.ArticleService
}

func NewArticleHandler(service services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: service}
}

func RegisterArticleHandlers(app *fiber.App, articleService services.ArticleService, jwtMiddleware fiber.Handler) {
	articleHandler := NewArticleHandler(articleService)

	app.Get("/articles", articleHandler.GetArticles)
	app.Post("/articles", jwtMiddleware, articleHandler.CreateArticle)
	app.Get("/articles/:id", articleHandler.GetArticle)
	app.Put("/articles/:id", jwtMiddleware, articleHandler.UpdateArticle)
	app.Delete("/articles/:id", jwtMiddleware, articleHandler.DeleteArticle)
}

func (h *ArticleHandler) GetArticles(c *fiber.Ctx) error {
	articles, err := h.articleService.GetArticles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(articles)
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	article := new(models.Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	createdArticle, err := h.articleService.CreateArticle(article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdArticle)
}

func (h *ArticleHandler) GetArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid article ID"})
	}

	article, err := h.articleService.GetArticle(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
	}
	return c.JSON(article)
}

func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid article ID"})
	}

	updates := new(models.Article)
	if err := c.BodyParser(updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON, ",
		})
	}
	guncellenenMakale, err := h.articleService.UpdateArticle(uint(id), updates)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Makale bulunamadı"})
	}

	return c.JSON(guncellenenMakale)
}
func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz makale ID"})
	}
	err = h.articleService.DeleteArticle(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Makale bulunamadı"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
