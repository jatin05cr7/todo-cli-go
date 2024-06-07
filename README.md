# Todo CLI in Go

A simple command-line to-do list application implemented in Go. This application allows users to add, remove, and mark tasks as completed. Tasks are saved to a file and loaded on startup for persistent storage.

## Features

- **Add Task:** Add a new task to the list.
- **Remove Task:** Remove an existing task by its ID.
- **Complete Task:** Mark a task as completed by its ID.
- **List Tasks:** Display all tasks with their completion status.
- **Persistent Storage:** Tasks are saved to a JSON file and loaded on startup.

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/jatin05cr7/todo-cli-go.git
   cd todo-cli-go

2.Initialize the Go module:
go mod init github.com/jatin05cr7/todo-cli-go

3.Build the application:
go build -o todo-cli main.go

Usage:-
Run the application:

Example Usage:-

To-Do List Application
1. Add Task
2. Remove Task
3. Complete Task
4. List Tasks
5. Exit
Enter your choice: 1
Enter task title: Buy groceries
Task added: Buy groceries



