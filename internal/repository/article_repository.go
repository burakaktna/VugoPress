package repository

import (
	"github.com/burakaktna/VugoPress/internal/models"
	"github.com/jinzhu/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *models.Article) error
	GetArticles() ([]*models.Article, error)
	GetArticle(id uint) (*models.Article, error)
	UpdateArticle(id uint, updates *models.Article) (*models.Article, error)
	DeleteArticle(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) CreateArticle(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) GetArticles() ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.Offset(0).Limit(10).Find(&articles).Error
	return articles, err
}

func (r *articleRepository) GetArticle(id uint) (*models.Article, error) {
	var article models.Article
	err := r.db.Where("id = ?", id).First(&article).Error
	return &article, err
}

func (r *articleRepository) UpdateArticle(id uint, updates *models.Article) (*models.Article, error) {
	var article models.Article
	err := r.db.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	article.Title = updates.Title
	article.Content = updates.Content
	err = r.db.Save(&article).Error
	return &article, err
}

func (r *articleRepository) DeleteArticle(id uint) error {
	return r.db.Delete(&models.Article{}, "id = ?", id).Error
}
