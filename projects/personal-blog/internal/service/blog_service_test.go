package service

import (
	"testing"

	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestPersonalBlogService_CreateArticle(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewPersonalBlogService(repo)

	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := service.CreateArticle(article)
	require.NoError(t, err)
	require.NotZero(t, id)
}

func TestPersonalBlogService_GetArticleByID(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewPersonalBlogService(repo)

	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    1,
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := service.CreateArticle(article)
	require.NoError(t, err)

	retrievedArticle, err := service.GetArticleByID(id)
	require.NoError(t, err)
	require.Equal(t, article, retrievedArticle)
}

func TestPersonalBlogService_DeleteArticle(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewPersonalBlogService(repo)

	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    1,
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := service.CreateArticle(article)
	require.NoError(t, err)

	err = service.DeleteArticle(id)
	require.NoError(t, err)
}

func TestPersonalBlogService_GetAllArticles(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewPersonalBlogService(repo)

	article1 := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    1,
			Title: "Test Article 1",
		},
		Content: "This is a test article.",
	}

	article2 := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    2,
			Title: "Test Article 2",
		},
		Content: "This is another test article.",
	}

	_, err := service.CreateArticle(article1)
	require.NoError(t, err)

	_, err = service.CreateArticle(article2)
	require.NoError(t, err)

	articles, err := service.GetAllArticlesSummary()
	require.NoError(t, err)
	require.Len(t, articles, 2)
	require.Equal(t, article1, articles[0])
	require.Equal(t, article2, articles[1])
}

func TestPersonalBlogService_UpdateArticle(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewPersonalBlogService(repo)

	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    1,
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := service.CreateArticle(article)
	require.NoError(t, err)

	updatedArticle := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    id,
			Title: "Updated Test Article",
		},
		Content: "This is an updated test article.",
	}

	err = service.UpdateArticle(updatedArticle)
	require.NoError(t, err)
}
