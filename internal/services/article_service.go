package services

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type ArticleService interface {
	Create(article *models.Article) (*models.Article, error)
	Index() ([]*models.Article, error)
	Show(id uint) (*models.Article, error)
	Update(id uint, updates *models.Article) (*models.Article, error)
	Delete(id uint) error
}

type articleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{db: db}
}

func (s *articleService) Create(article *models.Article) (*models.Article, error) {
	err := s.db.Create(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *articleService) Index() ([]*models.Article, error) {
	var articles []*models.Article
	err := s.db.Find(&articles).Error
	return articles, err
}

func (s *articleService) Show(id uint) (*models.Article, error) {
	var article models.Article
	err := s.db.First(&article, id).Error
	return &article, err
}

func (s *articleService) Update(id uint, updates *models.Article) (*models.Article, error) {
	var article models.Article
	if err := s.db.First(&article, id).Error; err != nil {
		return nil, err
	}
	err := s.db.Model(&article).Updates(updates).Error
	return &article, err
}

func (s *articleService) Delete(id uint) error {
	return s.db.Delete(&models.Article{}, id).Error
}
