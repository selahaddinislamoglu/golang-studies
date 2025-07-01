package service

import "github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"

type Service interface {
	CreateArticle(article *model.Article) (int64, error)
	UpdateArticle(article *model.Article) error
	DeleteArticle(id int64) error
	GetArticleByID(id int64) (*model.Article, error)
	GetAllArticlesSummary() ([]*model.ArticleSummary, error)
}
