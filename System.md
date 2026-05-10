# Student Task Manager System/

## Problem Statement

Students often struggle with organizing and keeping track of their academic tasks such as assignments, projects, quizzes, and deadlines. Because of this, some tasks may be forgotten or completed late.

The Student Task Manager is a simple Command-Line Interface (CLI) application developed using the Go programming language that helps students manage their daily tasks efficiently. The system allows users to add tasks, view saved tasks, mark tasks as completed, and delete unnecessary tasks.

The application is lightweight, easy to use, and designed to demonstrate the core features and capabilities of Golang.

---

# Functional Requirements

- [ ] Users can add new tasks
- [ ] Users can view all saved tasks
- [ ] Users can mark tasks as completed
- [ ] Users can delete tasks
- [ ] The system displays task status (Completed or Pending)
- [ ] The system stores tasks locally using a text or JSON file
- [ ] The system automatically loads saved tasks when the application starts

---

# Non-Functional Requirements

- The application must be built using Golang
- The application must run in the command line interface (CLI)
- The system must store data locally using a JSON or text file
- The application must run locally on the user's computer
- The system should respond quickly and smoothly
- Code should be simple, readable, and modular
- The application should handle invalid user inputs properly

---

# Architecture & Tech Stack

## Tech Stack

| Layer | Choice |
|---|---|
| Language | Golang |
| Interface | Command-Line Interface (CLI) |
| Storage | JSON File |
| Concurrency | Goroutines |
| Input Handling | bufio package |

---

## Architecture Overview

The application follows a simple CLI-based architecture where the program processes user input, performs task operations, and stores task data locally.

The system is divided into several components to separate responsibilities:

- **Main Layer**: Handles program execution and menu navigation
- **Task Management Layer**: Handles adding, viewing, completing, and deleting tasks
- **Storage Layer**: Manages saving and loading task data from a JSON file
- **Utility Layer**: Contains helper functions and reusable logic

---

## Project Structure

```bash
project/
├── main.go              # Entry point of the application
├── models/
│   └── task.go          # Task structure/model
├── services/
│   └── taskService.go   # Task management logic
├── data/
│   └── tasks.json       # Stores saved tasks
├── utils/
│   └── helper.go        # Helper functions
├── go.mod               # Go module file
└── README.md            # Project documentation