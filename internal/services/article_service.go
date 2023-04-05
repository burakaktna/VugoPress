package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/burakaktna/VugoPress/internal/repository"
)

type ArticleService interface {
	CreateArticle(article *models.Article) (*models.Article, error)
	GetArticles() ([]*models.Article, error)
	GetArticle(id uint) (*models.Article, error)
	UpdateArticle(id uint, updates *models.Article) (*models.Article, error)
	DeleteArticle(id uint) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(article *models.Article) (*models.Article, error) {
	err := s.repo.CreateArticle(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *articleService) GetArticles() ([]*models.Article, error) {
	return s.repo.GetArticles()
}

func (s *articleService) GetArticle(id uint) (*models.Article, error) {
	return s.repo.GetArticle(id)
}

func (s *articleService) UpdateArticle(id uint, updates *models.Article) (*models.Article, error) {
	return s.repo.UpdateArticle(id, updates)
}

func (s *articleService) DeleteArticle(id uint) error {
	return s.repo.DeleteArticle(id)
}
