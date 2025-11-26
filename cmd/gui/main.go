package main

import (
	"TaskTracker/internal/crud"
	"TaskTracker/internal/storage"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func filterTasks(tasks []storage.Task, filter string) []storage.Task {
	if filter == "All" || filter == "" {
		return tasks
	}
	var result []storage.Task
	for _, t := range tasks {
		switch filter {
		case "To Do":
			if t.Status == crud.StatusTodo {
				result = append(result, t)
			}
		case "In Progress":
			if t.Status == crud.StatusInProgress {
				result = append(result, t)
			}
		case "Done":
			if t.Status == crud.StatusDone {
				result = append(result, t)
			}
		}
	}
	return result
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Планировщик задач")
	myWindow.Resize(fyne.NewSize(900, 600))

	tasks, err := storage.LoadTasks()
	if err != nil {
		panic(fmt.Errorf("failed to load tasks: %w", err))
	}

	var taskList *widget.List
	var currentFilter = "All"

	entry := widget.NewEntry()
	entry.PlaceHolder = "Введите новый текст"

	addButton := widget.NewButton("Добавить задачу", func() {
		desc := entry.Text
		if desc == "" {
			return
		}
		updatedTasks, err := crud.AddTask(tasks, desc)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		tasks = updatedTasks
		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}
		entry.SetText("")
		taskList.Refresh()
	})

	filterSelect := widget.NewSelect([]string{"All", "To Do", "In Progress", "Done"}, func(value string) {
		currentFilter = value
		taskList.Refresh()
	})

	taskList = widget.NewList(
		func() int {
			return len(filterTasks(tasks, currentFilter))
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("ID"),
				widget.NewLabel("Description"),
				widget.NewLabel("Status"),
				widget.NewLabel("Created"),
				widget.NewButton("🗑️", nil),
				widget.NewButton("🔄", nil),
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			filtered := filterTasks(tasks, currentFilter)
			t := filtered[id]
			row := item.(*fyne.Container)

			idLabel := row.Objects[0].(*widget.Label)
			descLabel := row.Objects[1].(*widget.Label)
			statusLabel := row.Objects[2].(*widget.Label)
			createdLabel := row.Objects[3].(*widget.Label)
			deleteBtn := row.Objects[4].(*widget.Button)
			toggleBtn := row.Objects[5].(*widget.Button)

			idLabel.SetText(fmt.Sprintf("%d", t.ID))
			descLabel.SetText(t.Description)
			statusLabel.SetText(t.Status)
			createdLabel.SetText(t.CreatedAt.Format("2006-01-02 15:04"))

			deleteBtn.OnTapped = func() {
				updatedTasks, err := crud.DeleteTask(tasks, t.ID)
				if err != nil {
					fmt.Printf("Error deleting task: %v\n", err)
					return
				}
				tasks = updatedTasks
				if err := storage.SaveTasks(tasks); err != nil {
					fmt.Printf("Error saving tasks: %v\n", err)
					return
				}
				taskList.Refresh()
			}

			toggleBtn.OnTapped = func() {
				var updatedTasks []storage.Task
				var err error

				switch t.Status {
				case crud.StatusTodo:
					updatedTasks, err = crud.MarkInProgress(tasks, t.ID)
				case crud.StatusInProgress:
					updatedTasks, err = crud.MarkDone(tasks, t.ID)
				case crud.StatusDone:
					updatedTasks, err = crud.MarkTodo(tasks, t.ID)
				default:
					return
				}

				if err != nil {
					fmt.Printf("Error updating status: %v\n", err)
					return
				}
				tasks = updatedTasks
				if err := storage.SaveTasks(tasks); err != nil {
					fmt.Printf("Error saving tasks: %v\n", err)
					return
				}
				taskList.Refresh()
			}
		},
	)
	// делаю поле для ввода равную ширине окна
	topBar := container.NewBorder(nil, nil, nil, addButton, entry)
	filterBar := container.NewHBox(widget.NewLabel("Фильтр:"), filterSelect)

	//чтобы добавленные задачи растягивались на всю длину окна
	content := container.NewBorder(
		container.NewVBox(topBar, filterBar),
		nil,
		nil,
		nil,
		taskList,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
