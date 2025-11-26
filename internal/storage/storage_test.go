package storage

import (
	"os"
	"testing"
)

func TestSaveLoadTasks(t *testing.T) {
	// Создаём тестовые задачи
	tasks := []Task{
		{ID: 1, Description: "Test task", Status: "todo"},
	}

	// Сохраняем
	err := SaveTasks(tasks)
	if err != nil {
		t.Fatalf("SaveTasks failed: %v", err)
	}
	defer os.Remove("tasks.json") // удаляем после теста

	// Загружаем
	loadedTasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}

	if len(loadedTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(loadedTasks))
	}
	if loadedTasks[0].Description != "Test task" {
		t.Errorf("Expected 'Test task', got '%s'", loadedTasks[0].Description)
	}
}
