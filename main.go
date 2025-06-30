package main

import (
	"tasktracercli/commands"
	"tasktracercli/models"
	"tasktracercli/storage"
	"tasktracercli/utils"
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleCommand(command []string, file *os.File) error {
	if len(command) == 0 {
		return fmt.Errorf("empty command")
	}

	cmd := command[0]
	switch cmd {
	case "add":
		{
			if len(command) < 2 {
				return errors.New("missing task description")
			}
			description := strings.Trim(strings.Join(command[1:], " "), `"`)

			NewTask := models.Task{
				ID:          0,
				Description: description,
				Status:      models.StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			if err := storage.AppendJson(file, &NewTask); err != nil {
				return fmt.Errorf("add task error: %v", err)
			}
			fmt.Printf("\nTask added successfully (ID: %d)\n", NewTask.ID)
		}
	case "update":
		{
			if len(command) < 3 {
				return fmt.Errorf("missing task description")
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}
			description := strings.Trim(strings.Join(command[2:], " "), `"`)
			if err := commands.UpdateTask(file, id, description, false); err != nil { // false - update description
				return fmt.Errorf("update task error: %w", err)
			}
			fmt.Printf("\nTask updated successfully (ID: %d)\n", id)
		}
	case "delete":
		{
			if len(command) < 2 {
				return fmt.Errorf("missing task ID")
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}
			if err := commands.DeleteTask(file, id); err != nil {
				return fmt.Errorf("error: %v", err)
			}
			fmt.Printf("\nTask deleted successfully (ID: %d)\n", id)
		}
	case "list":
		{
			statusFilter := "all"
			if len(command) > 1 {
				statusFilter = command[1]
			}

			if err := commands.PrintTasks(file, statusFilter); err != nil {
				return fmt.Errorf("list tasks error: %w", err)
			}
		}

	case "mark-in-progress", "mark-done":
		{
			if len(command) < 2 {
				return errors.New("missing task ID")
			}
			status := strings.Join(strings.Split(cmd, "-")[1:], " ")
			id, err := strconv.Atoi(command[1])
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}

			if err := commands.UpdateTask(file, id, status, true); err != nil { // true - update status
				return fmt.Errorf("error: %v", err)
			}
			fmt.Printf("\nTask status successfully updated(ID: %d)\n", id)
		}
	default:
		return fmt.Errorf("unknown command: %s (type 'help' for available commands)", cmd)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	color.Cyan("Welcome to Task Tracker CLI!")
	fmt.Println("Type 'help' for available commands, 'exit' to quit\n")

	for {
		color.New(color.FgGreen).Print("task-cli> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "":
			continue
		case "exit":
			color.Yellow("Goodbye!")
			return
		case "help":
			utils.ShowHelp()
			continue
		}

		command := strings.Fields(input)
		file, err := os.OpenFile(commands.FileName, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			color.Red("Error: %v", err)
			continue
		}
		defer file.Close()

		if err := handleCommand(command, file); err != nil {
			color.Red("Error: %v", err)
		}

	}

}
