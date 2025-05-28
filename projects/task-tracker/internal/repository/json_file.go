package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	task.CreatedAt = time.Now().Format(time.RFC3339)
	task.UpdatedAt = time.Now().Format(time.RFC3339)

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

func (r *FileRepository) GetTaskByID(id int) (*model.Task, error) {
	for _, task := range *r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (r *FileRepository) DeleteTask(task *model.Task) error {
	for i, t := range *r.tasks {
		if t.ID == task.ID {
			*r.tasks = append((*r.tasks)[:i], (*r.tasks)[i+1:]...)
			return r.writeTasks()
		}
	}
	return fmt.Errorf("task with ID %d not found", task.ID)
}

func (r *FileRepository) UpdateTask(task *model.Task) error {
	for i, t := range *r.tasks {
		if t.ID == task.ID {
			task.UpdatedAt = time.Now().Format(time.RFC3339)
			(*r.tasks)[i] = *task
			return r.writeTasks()
		}
	}
	return fmt.Errorf("task with ID %d not found", task.ID)
}
