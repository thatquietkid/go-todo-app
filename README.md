# TodoApp - CLI ToDo Application in Go

TodoApp is a simple, command-line based ToDo application built in Go. It allows users to manage a list of tasks via intuitive terminal commands. Tasks are stored in a `tasks.csv` file, and the app supports adding, listing, completing, and deleting tasks.

## 📁 Project Structure
```
root/
├── go.mod
├── go.sum
├── main.go
├── tasks.csv
├── LICENSE
├── README.md
├── cmd/
│   ├── add.go         # Command to add a new task
│   ├── complete.go    # Command to mark a task as completed
│   ├── delete.go      # Command to delete a task by ID
│   ├── list.go        # Command to list all tasks
│   └── root.go        # Root command setup
└── internal/app/
├── constants.go   # App-wide constants
└── io.go          # CSV read/write logic
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.2 or later
- Git (for cloning the repo)

### Installation

```bash
git clone https://github.com/thatquietkid/go-todo-app.git
cd go-todo-app
go build -o todo
./todo
```

## 📌 Usage

### Root Command

```bash
./todo
```

This shows a list of available subcommands.

---

### Add a Task

```bash
./todo add [title] [description] [due-date] [completed]
```

* `title`: Title of the task
* `description`: Details about the task
* `due-date`: Due date in `YYYY-MM-DD` format
* `completed`: `true` or `false`

**Example:**

```bash
./todo add "Buy milk" "Get it from the nearby store" 2025-06-01 false
```

---

### List All Tasks

```bash
./todo list
```

Displays all tasks in a tabular format, including time left until the due date.

---

### Complete a Task

```bash
./todo complete [taskID]
```

Marks the task with the specified ID as completed.

**Aliases:** `done`

**Example:**

```bash
./todo complete 2
```

---

### Delete a Task

```bash
./todo delete [ID]
```

Deletes the task with the given ID and reassigns task IDs sequentially.

**Example:**

```bash
./todo delete 3
```

---

## 📄 Data Storage

All tasks are stored in a CSV file named `tasks.csv` located at the root of the project. This file must not be manually edited unless you know what you are doing.

---

## 🛠 Technologies Used

* Go
* Cobra CLI framework
* Standard library for CSV and time manipulation

---

## 📃 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## 🙌 Acknowledgments

* [spf13/cobra](https://github.com/spf13/cobra) for the CLI framework
* [mergestat/timediff](https://github.com/mergestat/timediff) for human-readable time differences

```

---
