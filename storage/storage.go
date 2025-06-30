package storage

import (
	"tasktracercli/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJson(file *os.File) ([]models.Task, error) {
	var tasks []models.Task

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("file stat error: %w", err)
	}

	if fileInfo.Size() == 0 {
		return tasks, nil
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return tasks, fmt.Errorf("seek error: %v", err)
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err != io.EOF {
		return tasks, fmt.Errorf("decode error: %v", err)
	}
	return tasks, nil
}

func WriteJson(file *os.File, tasks []models.Task) error {
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("truncate error: %v", err)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("seek error: %v", err)
	}

	if err := file.Sync(); err != nil {
		return fmt.Errorf("sync error: %v", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("encode error :%v", err)
	}
	return nil
}

func AppendJson(file *os.File, task *models.Task) error {
	tasks, err := ReadJson(file)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	var lastID int
	if len(tasks) != 0 {
		lastID = tasks[len(tasks)-1].ID
	}
	task.ID = lastID + 1

	tasks = append(tasks, *task)
	return WriteJson(file, tasks)
}
