package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
)

type FileRepository struct {
	storageDir string
	nextID     int64
}

func NewFileRepository(storageDir string) (Repository, error) {

	fileInfo, err := os.Stat(storageDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(storageDir, 0755)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else if !fileInfo.IsDir() {
		return nil, os.ErrInvalid
	}
	return &FileRepository{
		storageDir: storageDir,
		nextID:     1,
	}, nil
}

func (r *FileRepository) getNextID() int64 {
	if r.nextID > 0 {
		id := r.nextID
		r.nextID++
		return id
	}
	files, _ := os.ReadDir(r.storageDir)
	if len(files) == 0 {
		r.nextID = 1
		return r.nextID
	}
	var maxID int64
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var id int64
		_, err := fmt.Sscanf(file.Name(), "%d", &id)
		if err == nil && id > maxID {
			maxID = id
		}
	}
	r.nextID = maxID + 1
	return r.nextID
}

func (r *FileRepository) CreateArticle(article *model.Article) (int64, error) {
	id := r.getNextID()
	osFile, err := os.Create(fmt.Sprintf("%s/%d", r.storageDir, id))
	if err != nil {
		return 0, fmt.Errorf("failed to create article file: %w", err)
	}
	defer osFile.Close()
	article.ID = id

	if err := json.NewEncoder(osFile).Encode(article); err != nil {
		return 0, fmt.Errorf("failed to write article to file: %w", err)
	}
	return id, nil
}

func (r *FileRepository) UpdateArticle(article *model.Article) error {
	osFile, err := os.OpenFile(fmt.Sprintf("%s/%d", r.storageDir, article.ID), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("failed to open article file: %w", err)
	}
	defer osFile.Close()

	if err := json.NewEncoder(osFile).Encode(article); err != nil {
		return fmt.Errorf("failed to write article to file: %w", err)
	}
	return nil
}

func (r *FileRepository) DeleteArticle(id int64) error {
	filePath := fmt.Sprintf("%s/%d", r.storageDir, id)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("article with ID %d does not exist", id)
	}
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete article file: %w", err)
	}
	return nil
}

func (r *FileRepository) GetArticleByID(id int64) (*model.Article, error) {
	filePath := fmt.Sprintf("%s/%d", r.storageDir, id)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open article file: %w", err)
	}
	defer file.Close()

	var article model.Article
	if err := json.NewDecoder(file).Decode(&article); err != nil {
		return nil, fmt.Errorf("failed to decode article file: %w", err)
	}
	return &article, nil
}

func (r *FileRepository) GetAllArticlesSummary() ([]*model.ArticleSummary, error) {
	files, err := os.ReadDir(r.storageDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read articles directory: %w", err)
	}

	var articles []*model.ArticleSummary
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := fmt.Sprintf("%s/%s", r.storageDir, file.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open article file: %w", err)
		}
		defer file.Close()

		var article model.ArticleSummary
		if err := json.NewDecoder(file).Decode(&article); err != nil {
			return nil, fmt.Errorf("failed to decode article file: %w", err)
		}
		articles = append(articles, &article)
	}
	return articles, nil
}
