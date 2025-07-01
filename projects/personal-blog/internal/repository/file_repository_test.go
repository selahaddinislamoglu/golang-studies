package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
	"github.com/stretchr/testify/require"
)

func TestFileRepository_CreateArticle(t *testing.T) {
	repo, err := NewFileRepository("./test_articles/")
	require.NoError(t, err)
	defer os.RemoveAll("./test_articles/")

	article := model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := repo.CreateArticle(&article)
	require.NoError(t, err)
	require.NotZero(t, id)

	filePath := fmt.Sprintf("./test_articles/%d", id)
	file, err := os.Open(filePath)
	require.NoError(t, err)
	defer file.Close()

	var savedArticle model.Article
	err = json.NewDecoder(file).Decode(&savedArticle)
	require.NoError(t, err)
	require.Equal(t, article.Title, savedArticle.Title)
	require.Equal(t, article.Content, savedArticle.Content)
}

func TestFileRepository_UpdateArticle(t *testing.T) {
	repo, err := NewFileRepository("./test_articles/")
	require.NoError(t, err)
	defer os.RemoveAll("./test_articles/")

	article := model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := repo.CreateArticle(&article)
	require.NoError(t, err)

	article.ID = id
	article.Title = "Updated Test Article"
	err = repo.UpdateArticle(&article)
	require.NoError(t, err)

	filePath := fmt.Sprintf("./test_articles/%d", id)
	file, err := os.Open(filePath)
	require.NoError(t, err)
	defer file.Close()

	var updatedArticle model.Article
	err = json.NewDecoder(file).Decode(&updatedArticle)
	require.NoError(t, err)
	require.Equal(t, article.Title, updatedArticle.Title)
}

func TestFileRepository_DeleteArticle(t *testing.T) {
	repo, err := NewFileRepository("./test_articles/")
	require.NoError(t, err)
	defer os.RemoveAll("./test_articles/")

	article := model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := repo.CreateArticle(&article)
	require.NoError(t, err)

	err = repo.DeleteArticle(id)
	require.NoError(t, err)

	filePath := fmt.Sprintf("./test_articles/%d", id)
	_, err = os.Stat(filePath)
	require.True(t, os.IsNotExist(err), "File should be deleted")
}

func TestFileRepository_GetArticleByID(t *testing.T) {
	repo, err := NewFileRepository("./test_articles/")
	require.NoError(t, err)
	defer os.RemoveAll("./test_articles/")

	article := model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: "Test Article",
		},
		Content: "This is a test article.",
	}

	id, err := repo.CreateArticle(&article)
	require.NoError(t, err)

	retrievedArticle, err := repo.GetArticleByID(id)
	require.NoError(t, err)
	require.Equal(t, article.Title, retrievedArticle.Title)
	require.Equal(t, article.Content, retrievedArticle.Content)
}

func TestFileRepository_GetAllArticles(t *testing.T) {
	repo, err := NewFileRepository("./test_articles/")
	require.NoError(t, err)
	defer os.RemoveAll("./test_articles/")

	articles := []model.Article{
		{ArticleSummary: model.ArticleSummary{

			Title: "Test Article 1",
		},
			Content: "This is the first test article.",
		},
		{ArticleSummary: model.ArticleSummary{

			Title: "Test Article 2",
		},
			Content: "This is the second test article.",
		},
	}

	for _, article := range articles {
		_, err := repo.CreateArticle(&article)
		require.NoError(t, err)
	}

	retrievedArticles, err := repo.GetAllArticlesSummary()
	require.NoError(t, err)
	require.Equal(t, len(articles), len(retrievedArticles))
	for i := range articles {
		require.Equal(t, articles[i].Title, retrievedArticles[i].Title)
	}
}
