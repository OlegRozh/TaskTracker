package crud

import (
	"errors"
	"fmt"
	"time"

	"TaskTracker/internal/storage"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

var (
	ErrEmptyLine    = errors.New("description is required")
	ErrTaskNotFound = errors.New("task not found")
)

func AddTask(task []storage.Task, description string) ([]storage.Task, error) {
	if description == "" {
		return nil, ErrEmptyLine
	}
	maxID := 0
	for _, t := range task {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	NewID := maxID + 1
	newTask := storage.Task{
		ID:          NewID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	updatedTasks := append(task, newTask)
	return updatedTasks, nil
}

func UpdateTask(task []storage.Task, id int, description string) ([]storage.Task, error) {
	if description == "" {
		return nil, ErrEmptyLine
	}
	for i, t := range task {
		if t.ID == id {
			task[i].Description = description
			task[i].UpdatedAt = time.Now()
			return task, nil
		}
	}
	return nil, ErrTaskNotFound
}

func DeleteTask(task []storage.Task, id int) ([]storage.Task, error) {
	for i, t := range task {
		if t.ID == id {
			task = append(task[:i], task[i+1:]...)
			return task, nil
		}
	}
	return nil, ErrTaskNotFound
}

func GetTask(task []storage.Task, status string) []storage.Task {
	if status == "" {
		return task
	}
	var filteredTasks []storage.Task
	for _, t := range task {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

func MarkInProgress(task []storage.Task, id int) ([]storage.Task, error) {
	for i, t := range task {
		if t.ID == id {
			task[i].Status = StatusInProgress
			task[i].UpdatedAt = time.Now()
			return task, nil
		}
	}
	return nil, fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
}

func MarkDone(tasks []storage.Task, id int) ([]storage.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = StatusDone
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
}

func MarkTodo(task []storage.Task, id int) ([]storage.Task, error) {
	for i, t := range task {
		if t.ID == id {
			task[i].Status = StatusTodo
			task[i].UpdatedAt = time.Now()
			return task, nil
		}
	}
	return nil, fmt.Errorf("%w: ID %d", ErrTaskNotFound, id)
}
