package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/selahaddinislamoglu/golang-studies/projects/task-tracker/internal/model"
)

type FileRepository struct {
	filePath string
	tasks    *[]model.Task
}

func NewFileRepository(filePath string) (*FileRepository, error) {

	fileRepo := &FileRepository{
		filePath: filePath,
		tasks:    &[]model.Task{},
	}

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fileRepo, nil
	}

	err = fileRepo.loadTasks()
	if err != nil {
		return nil, err
	}
	return fileRepo, nil
}

func (r *FileRepository) loadTasks() error {

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	r.tasks = &[]model.Task{}

	err = json.Unmarshal(data, &r.tasks)
	if err != nil {
		return err
	}
	return nil
}

func (r *FileRepository) writeTasks() error {
	data, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file: %w", err)
	}
	return nil
}

func (r *FileRepository) AddTask(task *model.Task) error {

	id := 1
	if len(*r.tasks) > 0 {
		id = (*r.tasks)[len(*r.tasks)-1].ID + 1
	}
	task.ID = id

	*r.tasks = append(*r.tasks, *task)

	err := r.writeTasks()
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	return nil
}

func (r *FileRepository) GetTaskByDescription(description string) *model.Task {
	for _, task := range *r.tasks {
		if task.Description == description {
			return &task
		}
	}
	return nil
}
