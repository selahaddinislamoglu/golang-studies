package service

import (
	"errors"
	"time"

	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/repository"
)

type PersonalBlogService struct {
	repository repository.Repository
}

func NewPersonalBlogService(repository repository.Repository) *PersonalBlogService {
	return &PersonalBlogService{
		repository: repository,
	}
}

func (s *PersonalBlogService) CreateArticle(article *model.Article) (int64, error) {
	if article.Title == "" || article.Content == "" {
		return 0, errors.New("invalid article")
	}
	article.UpdatedAt = time.Now().UTC().Format(time.DateOnly)
	return s.repository.CreateArticle(article)
}

func (s *PersonalBlogService) UpdateArticle(article *model.Article) error {
	if article.Title == "" || article.Content == "" || article.ID == 0 {
		return errors.New("invalid article")
	}
	article.UpdatedAt = time.Now().UTC().Format(time.DateOnly)
	return s.repository.UpdateArticle(article)
}

func (s *PersonalBlogService) DeleteArticle(id int64) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	if _, err := s.GetArticleByID(id); err != nil {
		return err
	}
	return s.repository.DeleteArticle(id)
}

func (s *PersonalBlogService) GetArticleByID(id int64) (*model.Article, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repository.GetArticleByID(id)
}

func (s *PersonalBlogService) GetAllArticlesSummary() ([]*model.ArticleSummary, error) {
	return s.repository.GetAllArticlesSummary()
}
