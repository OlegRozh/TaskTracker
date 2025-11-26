# TaskTracker — Console and Graphical Task Manager

## Description

TaskTracker is an application for managing tasks with **two operation modes**:

- **Console Interface (CLI)** — for working in the terminal.
- **Graphical Interface (GUI)** — with a visual interface powered by Fyne.

## Key Features

- Adding new tasks
- Editing task descriptions
- Deleting tasks
- Changing task status:
    - `todo` (to do)
    - `in-progress` (in progress)
    - `done` (completed)
- Filtering tasks by status
- Saving all data to a JSON file

## Installation and Launch

### Option A: Graphical Interface (GUI)

Open file cmd/gui/TaskTracker.exe 
or go run cmd/gui/main.go(is required Go 1.20+ and Fyne 2.4.4+)

### Option B: Console Interface (CLI)

go run cmd/cli/main.go

#### Adding a new task
task-cli add "Buy groceries"
Output: Task added successfully
#### Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
#### Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1
#### Listing all tasks
task-cli list
#### Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress