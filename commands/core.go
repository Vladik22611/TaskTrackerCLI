package commands

import (
	"tasktracercli/models"
	"tasktracercli/storage"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	FileName         string = "data.json"
	ErrTaskNotFound         = errors.New("task not found")
	ErrInvalidStatus        = errors.New("invalid status")
)

func isValidStatus(status string) bool {
	switch status {
	case models.StatusTodo, models.StatusInProgress, models.StatusDone:
		return true
	}
	return false
}

func UpdateTask(file *os.File, id int, update_element string, isStatusUpdate bool) error {
	if isStatusUpdate && !isValidStatus(update_element) {
		return ErrInvalidStatus
	}
	tasks, err := storage.ReadJson(file)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	for ind, task := range tasks {
		if task.ID == id {
			if isStatusUpdate {
				tasks[ind].Status = update_element
			} else {
				tasks[ind].Description = update_element
			}
			return storage.WriteJson(file, tasks)
		}
	}
	return ErrTaskNotFound
}

func DeleteTask(file *os.File, id int) error {
	tasks, err := storage.ReadJson(file)
	if err != nil {
		return fmt.Errorf("read error : %w", err)
	}
	for ind, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:ind], tasks[ind+1:]...)
			return storage.WriteJson(file, tasks)
		}
	}
	return ErrTaskNotFound
}

func PrintTasks(file *os.File, statusFilter string) error {
	tasks, err := storage.ReadJson(file)
	if err != nil {
		return fmt.Errorf("read error %w", err)
	}

	if len(tasks) == 0 {
		return fmt.Errorf("task list is empty")
	}

	var count int
	for _, task := range tasks {
		if statusFilter == "all" || task.Status == statusFilter {
			count++
			printTask(task)
		}
	}
	if count == 0 {
		fmt.Printf("No tasks with status '%s'\n", statusFilter)
	}
	return nil
}

func printTask(task models.Task) {
	fmt.Printf("Task #%d:\n", task.ID)
	fmt.Printf("  ID:          %d\n", task.ID)
	fmt.Printf("  Description: %s\n", task.Description)

	// Цветной вывод статуса
	switch task.Status {
	case models.StatusTodo:
		color.Red("  Status:      %s\n", task.Status)
	case models.StatusInProgress:
		color.Yellow("  Status:      %s\n", task.Status)
	case models.StatusDone:
		color.Green("  Status:      %s\n", task.Status)
	default:
		fmt.Printf("  Status:      %s\n", task.Status)
	}

	fmt.Printf("  CreatedAt:   %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  UpdatedAt:   %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("--------------------------------------")
}
