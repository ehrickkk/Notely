// Package utils provides reusable helper functions for the Notely CLI application.
package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// PrintHeader prints the Notely application banner at the top of the screen.
func PrintHeader() {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║         NOTELY – Student Task Manager       ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()
}

// PrintDivider prints a horizontal line to visually separate sections.
func PrintDivider() {
	fmt.Println("  ──────────────────────────────────────────")
}

// ClearScreen clears the terminal by printing blank lines,
// giving the appearance of a fresh screen on all platforms.
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// ReadInput reads a full line of text from the user via the provided bufio.Reader.
// It trims any leading/trailing whitespace and newline characters before returning.
func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	// Trim \r\n (Windows) or \n (Unix) and surrounding spaces
	return strings.TrimSpace(input)
}

// PressEnterToContinue pauses execution and waits for the user to press Enter.
// Used after displaying results so the user has time to read them.
func PressEnterToContinue(reader *bufio.Reader) {
	fmt.Print("\n  Press Enter to continue...")
	reader.ReadString('\n')
}

// ValidateNotEmpty returns true if the given string is not blank after trimming.
// Used to ensure required fields are not left empty by the user.
func ValidateNotEmpty(value string) bool {
	return strings.TrimSpace(value) != ""
}
