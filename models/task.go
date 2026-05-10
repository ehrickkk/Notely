// Package models defines the data structures used in the Notely application.
package models

import "time"

// Task represents a single student task with all its details.
// Each task is stored as a JSON object in the data/tasks.json file.
type Task struct {
	ID          int       `json:"id"`          // Unique identifier for the task
	Title       string    `json:"title"`       // Short name of the task
	Description string    `json:"description"` // Detailed description of what needs to be done
	Subject     string    `json:"subject"`     // School subject the task belongs to
	DueDate     string    `json:"due_date"`    // Deadline for the task (e.g., "2026-05-20")
	Status      string    `json:"status"`      // Current status: "Pending" or "Completed"
	CreatedAt   time.Time `json:"created_at"`  // Timestamp when the task was created
}
