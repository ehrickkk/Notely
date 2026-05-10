# Notely – Student Task Manager

A lightweight Command-Line Interface (CLI) application built with **Go (Golang)** to help students organize and track their academic tasks such as assignments, quizzes, and projects.

---

## Features

| Feature | Description |
|---|---|
| Add Task | Create a new task with title, description, subject, and due date |
| View All Tasks | Display all tasks with their full details and status |
| Mark as Completed | Update a task's status from Pending to Completed |
| Delete Task | Permanently remove a task after confirmation |
| Auto-Save | Tasks are auto-saved every 30 seconds via a background goroutine |
| Persistent Storage | All tasks are saved to `data/tasks.json` and reloaded on startup |

---

## Project Structure

```
Notely/
├── main.go                 # Entry point – main menu loop
├── models/
│   └── task.go             # Task struct / data model
├── services/
│   └── taskService.go      # All task business logic
├── utils/
│   └── helper.go           # Reusable helper functions (I/O, UI)
├── data/
│   └── tasks.json          # Local JSON storage for tasks
├── go.mod                  # Go module definition
└── README.md               # Project documentation
```

---

## Prerequisites

- [Go](https://go.dev/dl/) **1.21** or higher installed on your machine
- A terminal / command prompt

Verify your Go installation:
```bash
go version
```

---

## How to Run

### 1. Navigate to the project folder

```bash
cd path/to/Notely
```

### 2. Run the application directly

```bash
go run main.go
```

### 3. (Optional) Build an executable

```bash
# Windows
go build -o notely.exe .

# macOS / Linux
go build -o notely .
```

Then run the executable:

```bash
# Windows
.\notely.exe

# macOS / Linux
./notely
```

---

## Usage

When the application starts, you will see the main menu:

```
╔════════════════════════════════════════════╗
║         NOTELY – Student Task Manager       ║
╚════════════════════════════════════════════╝

  MAIN MENU
  ──────────────────────────────────────────
  [1]  Add Task
  [2]  View All Tasks
  [3]  Mark Task as Completed
  [4]  Delete Task
  [0]  Exit
  ──────────────────────────────────────────
  Enter your choice:
```

- Type a number (0–4) and press **Enter** to navigate.
- Follow the on-screen prompts for each action.
- Tasks are automatically saved to `data/tasks.json`.

---

## Data Storage

Tasks are stored in `data/tasks.json` in a human-readable format:

```json
[
  {
    "id": 1,
    "title": "Math Assignment",
    "description": "Complete exercises 1–20 on page 45",
    "subject": "Mathematics",
    "due_date": "2026-05-15",
    "status": "Pending",
    "created_at": "2026-05-11T08:00:00Z"
  }
]
```

---

## Tech Stack

| Component | Technology |
|---|---|
| Language | Go (Golang) |
| Interface | CLI (Command-Line Interface) |
| Storage | JSON File (`data/tasks.json`) |
| Input Handling | `bufio` package |
| Concurrency | Goroutines (background auto-save) |
| Data Encoding | `encoding/json` package |

---

## Authors

Built as a student project demonstrating core Golang concepts including structs, slices, file I/O, goroutines, and modular package design.

# Notely
A Student Task Manager System
