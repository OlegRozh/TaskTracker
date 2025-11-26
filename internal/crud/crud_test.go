// internal/crud/crud_test.go
package crud

import (
	"TaskTracker/internal/storage"
	"errors"
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks := make([]storage.Task, 0)
	description := "Buy groceries"

	updatedTasks, err := AddTask(tasks, description)

	if err != nil {
		t.Fatalf("AddTask returned an error: %v", err)
	}
	if len(updatedTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(updatedTasks))
	}
	if updatedTasks[0].Description != description {
		t.Errorf("Expected description '%s', got '%s'", description, updatedTasks[0].Description)
	}
	if updatedTasks[0].Status != StatusTodo {
		t.Errorf("Expected status '%s', got '%s'", StatusTodo, updatedTasks[0].Status)
	}
	if updatedTasks[0].ID != 1 {
		t.Errorf("Expected ID 1, got %d", updatedTasks[0].ID)
	}
}

func TestAddTask_EmptyDescription(t *testing.T) {
	tasks := make([]storage.Task, 0)
	_, err := AddTask(tasks, "")

	if err == nil {
		t.Error("Expected error for empty description, got nil")
	}
	if !errors.Is(err, ErrEmptyLine) {
		t.Errorf("Expected ErrEmptyLine, got %v", err)
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Old task", Status: StatusTodo},
	}
	newDescription := "Updated task"

	updatedTasks, err := UpdateTask(tasks, 1, newDescription)

	if err != nil {
		t.Fatalf("UpdateTask returned an error: %v", err)
	}
	if len(updatedTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(updatedTasks))
	}
	if updatedTasks[0].Description != newDescription {
		t.Errorf("Expected description '%s', got '%s'", newDescription, updatedTasks[0].Description)
	}
}

func TestUpdateTask_NotFound(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task", Status: StatusTodo},
	}
	_, err := UpdateTask(tasks, 999, "New description")

	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
	if !errors.Is(err, ErrTaskNotFound) { // ← исправлено
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}

func TestUpdateTask_EmptyDescription(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task", Status: StatusTodo},
	}
	_, err := UpdateTask(tasks, 1, "")

	if err == nil {
		t.Error("Expected error for empty description, got nil")
	}
	if !errors.Is(err, ErrEmptyLine) { // ← исправлено
		t.Errorf("Expected ErrEmptyLine, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task 1", Status: StatusTodo},
		{ID: 2, Description: "Task 2", Status: StatusTodo},
	}

	updatedTasks, err := DeleteTask(tasks, 1)

	if err != nil {
		t.Fatalf("DeleteTask returned an error: %v", err)
	}
	if len(updatedTasks) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(updatedTasks))
	}
	if updatedTasks[0].ID != 2 {
		t.Errorf("Expected remaining task ID 2, got %d", updatedTasks[0].ID)
	}
}

func TestDeleteTask_NotFound(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task", Status: StatusTodo},
	}
	_, err := DeleteTask(tasks, 999)

	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
	if !errors.Is(err, ErrTaskNotFound) { // ← исправлено
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}

func TestMarkInProgress(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task", Status: StatusTodo},
	}

	updatedTasks, err := MarkInProgress(tasks, 1)

	if err != nil {
		t.Fatalf("MarkInProgress returned an error: %v", err)
	}
	if updatedTasks[0].Status != StatusInProgress {
		t.Errorf("Expected status '%s', got '%s'", StatusInProgress, updatedTasks[0].Status)
	}
}

func TestMarkDone(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task", Status: StatusTodo},
	}

	updatedTasks, err := MarkDone(tasks, 1)

	if err != nil {
		t.Fatalf("MarkDone returned an error: %v", err)
	}
	if updatedTasks[0].Status != StatusDone {
		t.Errorf("Expected status '%s', got '%s'", StatusDone, updatedTasks[0].Status)
	}
}

func TestGetTask(t *testing.T) {
	tasks := []storage.Task{
		{ID: 1, Description: "Task 1", Status: StatusTodo},
		{ID: 2, Description: "Task 2", Status: StatusInProgress},
		{ID: 3, Description: "Task 3", Status: StatusDone},
	}

	// Тесты фильтрации
	todoTasks := GetTask(tasks, StatusTodo)
	if len(todoTasks) != 1 || todoTasks[0].ID != 1 {
		t.Errorf("Expected 1 todo task with ID 1, got %d tasks", len(todoTasks))
	}

	inProgressTasks := GetTask(tasks, StatusInProgress)
	if len(inProgressTasks) != 1 || inProgressTasks[0].ID != 2 {
		t.Errorf("Expected 1 in-progress task with ID 2, got %d tasks", len(inProgressTasks))
	}

	doneTasks := GetTask(tasks, StatusDone)
	if len(doneTasks) != 1 || doneTasks[0].ID != 3 {
		t.Errorf("Expected 1 done task with ID 3, got %d tasks", len(doneTasks))
	}

	allTasks := GetTask(tasks, "")
	if len(allTasks) != 3 {
		t.Errorf("Expected 3 tasks with empty filter, got %d", len(allTasks))
	}
}
