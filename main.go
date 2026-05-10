// Notely – Student Task Manager
// A simple CLI application to help students manage their academic tasks.
// Built with Go using structs, slices, file I/O, goroutines, and the bufio package.
package main

import (
	"bufio"
	"fmt"
	"notely/services"
	"notely/utils"
	"os"
	"time"
)

func main() {
	// Create a buffered reader for all user input throughout the application
	reader := bufio.NewReader(os.Stdin)

	// Initialise the task service – this automatically loads any previously saved tasks
	taskService := services.NewTaskService()

	// Start the background auto-save goroutine (saves every 30 seconds)
	// A channel is used to signal a clean shutdown when the user exits
	stopAutoSave := make(chan struct{})
	go taskService.AutoSave(30*time.Second, stopAutoSave)

	// ── Main application loop ──────────────────────────────────────────────────
	for {
		utils.ClearScreen()
		utils.PrintHeader()

		fmt.Println("  MAIN MENU")
		utils.PrintDivider()
		fmt.Println("  [1]  Add Task")
		fmt.Println("  [2]  View All Tasks")
		fmt.Println("  [3]  Mark Task as Completed")
		fmt.Println("  [4]  Delete Task")
		fmt.Println("  [0]  Exit")
		utils.PrintDivider()
		fmt.Print("  Enter your choice: ")

		choice := utils.ReadInput(reader)

		switch choice {
		case "1":
			taskService.AddTask(reader)

		case "2":
			taskService.ViewTasks(reader)

		case "3":
			taskService.MarkCompleted(reader)

		case "4":
			taskService.DeleteTask(reader)

		case "0":
			// Signal the auto-save goroutine to stop
			close(stopAutoSave)

			// Perform a final save to ensure no data is lost on exit
			taskService.SaveTasks()

			utils.ClearScreen()
			utils.PrintHeader()
			fmt.Println("  Thank you for using Notely!")
			fmt.Println("  All your tasks have been saved. See you next time!")
			fmt.Println()
			os.Exit(0)

		default:
			fmt.Println("\n  [!] Invalid choice. Please enter a number from 0 to 4.")
			utils.PressEnterToContinue(reader)
		}
	}
}
