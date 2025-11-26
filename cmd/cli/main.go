package main

import (
	"TaskTracker/internal/crud"
	"TaskTracker/internal/storage"
	"fmt"
	"os"
	"strconv"
)

func printTasks(tasks []storage.Task) {
	if len(tasks) == 0 {
		fmt.Println("tasks not found")
		return
	}
	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | [%s] | created: %s | updated: %s\n",
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format("2006-01-02 15:04:05"),
			t.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	command := os.Args[1]
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		description := os.Args[2]
		updatedTasks, err := crud.AddTask(tasks, description)
		if err != nil {
			fmt.Printf("Error adding tasks: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully ID:")
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		description := os.Args[3]
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer")
			os.Exit(1)
		}
		updatedTasks, err := crud.UpdateTask(tasks, id, description)
		if err != nil {
			fmt.Printf("Error adding tasks: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача успешно обновлена")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer")
			os.Exit(1)
		}
		updatedTasks, err := crud.DeleteTask(tasks, id)
		if err != nil {
			fmt.Printf("Error deleting tasks: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
	case "list":
		var status string
		if len(os.Args) > 2 {
			status = os.Args[2]
		} else {
			status = ""
		}
		filtered := crud.GetTask(tasks, status)
		printTasks(filtered)
		return
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer")
			os.Exit(1)
		}
		updatedTasks, err := crud.MarkInProgress(tasks, id)
		if err != nil {
			fmt.Printf("Error changing status: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task marked as in-progress")
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer")
			os.Exit(1)
		}
		updatedTasks, err := crud.MarkDone(tasks, id)
		if err != nil {
			fmt.Printf("Error changing status: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task marked as done")
	case "mark-todo":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be an integer")
			os.Exit(1)
		}
		updatedTasks, err := crud.MarkTodo(tasks, id)
		if err != nil {
			fmt.Printf("Error changing status: %v\n", err)
			os.Exit(1)
		}
		err = storage.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task marked is created")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
