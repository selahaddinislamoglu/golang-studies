package service

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
)

type MockPersonalBlogService struct {
	articles []*model.Article
}

func NewMockPersonalBlogService() *MockPersonalBlogService {
	return &MockPersonalBlogService{}
}

func (s *MockPersonalBlogService) GetAllArticles() ([]*model.Article, error) {
	return s.articles, nil
}

func (s *MockPersonalBlogService) CreateArticle(article *model.Article) (int64, error) {
	s.articles = append(s.articles, article)
	return int64(len(s.articles)), nil
}

func (s *MockPersonalBlogService) DeleteArticle(id int64) error {
	for i, article := range s.articles {
		if article.ID == id {
			s.articles = append(s.articles[:i], s.articles[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *MockPersonalBlogService) UpdateArticle(article *model.Article) error {
	for i, a := range s.articles {
		if a.ID == article.ID {
			s.articles[i] = article
			return nil
		}
	}
	return nil
}

func (s *MockPersonalBlogService) GetArticleByID(id int64) (*model.Article, error) {
	for _, article := range s.articles {
		if article.ID == id {
			return article, nil
		}
	}
	return nil, nil
}
