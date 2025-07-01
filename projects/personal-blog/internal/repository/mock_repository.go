package repository

import "github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"

type MockRepository struct {
	articles []*model.Article
}

func NewMockRepository() Repository {
	return &MockRepository{}
}

func (r *MockRepository) CreateArticle(article *model.Article) (int64, error) {
	r.articles = append(r.articles, article)
	return int64(len(r.articles)), nil
}

func (r *MockRepository) GetArticleByID(id int64) (*model.Article, error) {
	for _, article := range r.articles {
		if article.ID == id {
			return article, nil
		}
	}
	return nil, nil
}

func (r *MockRepository) GetAllArticles() ([]*model.Article, error) {
	return r.articles, nil
}

func (r *MockRepository) DeleteArticle(id int64) error {
	for i, article := range r.articles {
		if article.ID == id {
			r.articles = append(r.articles[:i], r.articles[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *MockRepository) UpdateArticle(article *model.Article) error {
	for i, a := range r.articles {
		if a.ID == article.ID {
			r.articles[i] = article
			return nil
		}
	}
	return nil
}

func (r *MockRepository) GetAllArticlesSummary() ([]*model.ArticleSummary, error) {
	var summaries []*model.ArticleSummary
	for _, article := range r.articles {
		summaries = append(summaries, &article.ArticleSummary)
	}
	return summaries, nil
}
