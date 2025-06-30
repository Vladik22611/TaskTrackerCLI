package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func ShowHelp() {
	// Create color printers
	header := color.New(color.FgCyan, color.Bold).PrintfFunc()
	command := color.New(color.FgGreen).PrintfFunc()
	example := color.New(color.FgYellow).PrintfFunc()
	param := color.New(color.FgMagenta).PrintfFunc()

	// Header
	header("\nTask Manager CLI - Help\n\n")

	// Core commands
	color.Blue("Basic Commands:\n")
	command("  add             ")
	param("<description>\t")
	fmt.Println("Add a new task")
	command("  list            ")
	param("[status]\t\t")
	fmt.Println("List tasks (optionally filtered by status)")
	command("  update          ")
	param("<ID> <description>\t")
	fmt.Println("Update task description")
	command("  delete          ")
	param("<ID>\t\t")
	fmt.Println("Delete a task")
	command("  mark-in-progress")
	param(" <ID>\t\t")
	fmt.Println("Change task status to 'in progress'")
	command("  mark-done       ")
	param("<ID>\t\t")
	fmt.Println("Mark task as completed")
	command("  exit            ")
	param("\t\t")
	fmt.Println("Exit the application\n")

	// Status filters
	color.Blue("Available Status Filters:\n")
	fmt.Println("  todo          - Show only pending tasks")
	fmt.Println("  in-progress   - Show active tasks")
	fmt.Println("  done          - Show completed tasks\n")

	// Examples
	color.Blue("Examples:\n")
	example("  $ ")
	fmt.Println("add \"Buy groceries\"")
	example("  $ ")
	fmt.Println("list in-progress")
	example("  $ ")
	fmt.Println("update 3 \"Buy organic milk\"")
	example("  $ ")
	fmt.Println("mark-done 3")
	example("  $ ")
	fmt.Println("delete 3\n")

}
