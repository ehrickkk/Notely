// Package services contains all the business logic for managing student tasks.
package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"notely/models"
	"notely/utils"
	"os"
	"strconv"
	"time"
)

// dataFilePath is the path to the JSON file where tasks are persisted.
const dataFilePath = "data/tasks.json"

// TaskService holds the in-memory list of tasks and provides all task operations.
type TaskService struct {
	Tasks []models.Task
}

// NewTaskService initialises a new TaskService and loads any previously saved tasks
// from the JSON file so the user's data is always available on startup.
func NewTaskService() *TaskService {
	ts := &TaskService{}
	ts.LoadTasks()
	return ts
}

// LoadTasks reads the JSON file and populates the in-memory task slice.
// If the file does not exist yet, it simply starts with an empty list.
func (ts *TaskService) LoadTasks() {
	file, err := os.ReadFile(dataFilePath)
	if err != nil {
		// File may not exist on first run – that is fine.
		ts.Tasks = []models.Task{}
		return
	}

	err = json.Unmarshal(file, &ts.Tasks)
	if err != nil {
		fmt.Println("  [!] Warning: Could not parse tasks file. Starting with an empty list.")
		ts.Tasks = []models.Task{}
	}
}

// SaveTasks writes the current in-memory task list to the JSON file.
// It creates the data/ directory if it does not already exist.
func (ts *TaskService) SaveTasks() error {
	// Ensure the data directory exists before writing
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return err
	}

	// Marshal with indentation for human-readable storage
	data, err := json.MarshalIndent(ts.Tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFilePath, data, 0644)
}

// getNextID returns an ID that is one higher than the highest existing task ID.
// This prevents ID collisions when new tasks are added.
func (ts *TaskService) getNextID() int {
	if len(ts.Tasks) == 0 {
		return 1
	}
	maxID := 0
	for _, task := range ts.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

// AddTask guides the user through entering a new task, validates required fields,
// appends it to the task list, and immediately saves it to disk.
func (ts *TaskService) AddTask(reader *bufio.Reader) {
	utils.ClearScreen()
	utils.PrintHeader()
	fmt.Println("  [ ADD NEW TASK ]")
	utils.PrintDivider()

	// --- Title (required) ---
	fmt.Print("  Task Title    : ")
	title := utils.ReadInput(reader)
	if !utils.ValidateNotEmpty(title) {
		fmt.Println("\n  [!] Title cannot be empty.")
		utils.PressEnterToContinue(reader)
		return
	}

	// --- Description (optional) ---
	fmt.Print("  Description   : ")
	description := utils.ReadInput(reader)

	// --- Subject (required) ---
	fmt.Print("  Subject       : ")
	subject := utils.ReadInput(reader)
	if !utils.ValidateNotEmpty(subject) {
		fmt.Println("\n  [!] Subject cannot be empty.")
		utils.PressEnterToContinue(reader)
		return
	}

	// --- Due Date (required) ---
	fmt.Print("  Due Date      : ")
	dueDate := utils.ReadInput(reader)
	if !utils.ValidateNotEmpty(dueDate) {
		fmt.Println("\n  [!] Due date cannot be empty.")
		utils.PressEnterToContinue(reader)
		return
	}

	// Build the new task and add it to the slice
	newTask := models.Task{
		ID:          ts.getNextID(),
		Title:       title,
		Description: description,
		Subject:     subject,
		DueDate:     dueDate,
		Status:      "Pending",
		CreatedAt:   time.Now(),
	}

	ts.Tasks = append(ts.Tasks, newTask)

	// Persist immediately so data is not lost if the program closes unexpectedly
	if err := ts.SaveTasks(); err != nil {
		fmt.Println("\n  [!] Warning: Could not save task:", err)
	} else {
		fmt.Println("\n  [✓] Task added and saved successfully!")
	}

	utils.PressEnterToContinue(reader)
}

// ViewTasks displays all tasks stored in memory with full details and status indicators.
func (ts *TaskService) ViewTasks(reader *bufio.Reader) {
	utils.ClearScreen()
	utils.PrintHeader()
	fmt.Println("  [ ALL TASKS ]")
	utils.PrintDivider()

	if len(ts.Tasks) == 0 {
		fmt.Println("  No tasks found. Start by adding a new task!")
	} else {
		for _, task := range ts.Tasks {
			// Visual indicator: filled circle = completed, empty circle = pending
			statusIcon := "○"
			if task.Status == "Completed" {
				statusIcon = "✓"
			}

			fmt.Printf("  [%s] #%d  %s\n", statusIcon, task.ID, task.Title)
			fmt.Printf("       Subject     : %s\n", task.Subject)
			fmt.Printf("       Description : %s\n", task.Description)
			fmt.Printf("       Due Date    : %s\n", task.DueDate)
			fmt.Printf("       Status      : %s\n", task.Status)
			fmt.Printf("       Created At  : %s\n", task.CreatedAt.Format("Jan 02, 2006  03:04 PM"))
			utils.PrintDivider()
		}
		fmt.Printf("  Total Tasks: %d\n", len(ts.Tasks))
	}

	utils.PressEnterToContinue(reader)
}

// MarkCompleted lists all pending tasks, prompts for an ID, and flips its status
// to "Completed", then saves the updated list to disk.
func (ts *TaskService) MarkCompleted(reader *bufio.Reader) {
	utils.ClearScreen()
	utils.PrintHeader()
	fmt.Println("  [ MARK TASK AS COMPLETED ]")
	utils.PrintDivider()

	if len(ts.Tasks) == 0 {
		fmt.Println("  No tasks available.")
		utils.PressEnterToContinue(reader)
		return
	}

	// Show only pending tasks so the user knows which ones still need attention
	fmt.Println("  Pending Tasks:")
	hasPending := false
	for _, task := range ts.Tasks {
		if task.Status == "Pending" {
			fmt.Printf("    #%d  %-30s  [%s]\n", task.ID, task.Title, task.Subject)
			hasPending = true
		}
	}

	if !hasPending {
		fmt.Println("\n  All tasks are already completed. Great job!")
		utils.PressEnterToContinue(reader)
		return
	}

	fmt.Print("\n  Enter Task ID to mark as completed: ")
	input := utils.ReadInput(reader)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\n  [!] Invalid input. Please enter a numeric Task ID.")
		utils.PressEnterToContinue(reader)
		return
	}

	found := false
	for i, task := range ts.Tasks {
		if task.ID == id {
			if task.Status == "Completed" {
				fmt.Printf("\n  [!] Task #%d is already marked as completed.\n", id)
			} else {
				ts.Tasks[i].Status = "Completed"
				if saveErr := ts.SaveTasks(); saveErr != nil {
					fmt.Println("\n  [!] Warning: Could not save changes:", saveErr)
				} else {
					fmt.Printf("\n  [✓] Task #%d \"%s\" has been marked as completed!\n", id, task.Title)
				}
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("\n  [!] No task found with ID #%d. Please check the ID and try again.\n", id)
	}

	utils.PressEnterToContinue(reader)
}

// DeleteTask lists all tasks, asks the user to confirm, then removes the chosen task
// from the slice and saves the updated list to disk.
func (ts *TaskService) DeleteTask(reader *bufio.Reader) {
	utils.ClearScreen()
	utils.PrintHeader()
	fmt.Println("  [ DELETE TASK ]")
	utils.PrintDivider()

	if len(ts.Tasks) == 0 {
		fmt.Println("  No tasks available to delete.")
		utils.PressEnterToContinue(reader)
		return
	}

	// Show all tasks so the user can pick the correct ID
	fmt.Println("  Available Tasks:")
	for _, task := range ts.Tasks {
		statusIcon := "○"
		if task.Status == "Completed" {
			statusIcon = "✓"
		}
		fmt.Printf("    [%s] #%d  %-30s  [%s]\n", statusIcon, task.ID, task.Title, task.Subject)
	}

	fmt.Print("\n  Enter Task ID to delete: ")
	input := utils.ReadInput(reader)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\n  [!] Invalid input. Please enter a numeric Task ID.")
		utils.PressEnterToContinue(reader)
		return
	}

	found := false
	for i, task := range ts.Tasks {
		if task.ID == id {
			// Ask for confirmation before permanently deleting
			fmt.Printf("\n  Are you sure you want to delete Task #%d \"%s\"? (y/n): ", task.ID, task.Title)
			confirm := utils.ReadInput(reader)

			if confirm == "y" || confirm == "Y" {
				// Remove the task by slicing it out of the list
				ts.Tasks = append(ts.Tasks[:i], ts.Tasks[i+1:]...)

				if saveErr := ts.SaveTasks(); saveErr != nil {
					fmt.Println("\n  [!] Warning: Could not save changes:", saveErr)
				} else {
					fmt.Printf("\n  [✓] Task #%d \"%s\" deleted successfully!\n", id, task.Title)
				}
			} else {
				fmt.Println("\n  Deletion cancelled.")
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("\n  [!] No task found with ID #%d. Please check the ID and try again.\n", id)
	}

	utils.PressEnterToContinue(reader)
}

// AutoSave runs in a goroutine and saves tasks to disk at a regular interval.
// It listens on stopCh so it can be shut down cleanly when the program exits.
func (ts *TaskService) AutoSave(interval time.Duration, stopCh <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Silently save in the background – no output to disrupt the UI
			ts.SaveTasks()
		case <-stopCh:
			// Program is exiting; stop the goroutine
			return
		}
	}
}
