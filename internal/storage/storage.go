package storage

import (
	"task-tracker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func ReadTasksFromFile() ([]model.Task, error) {
	filePath := "tasks.json"
	jsonFile, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err := os.Create(filePath)
			if err != nil {
				return nil, fmt.Errorf("error creating tasks.json file: %w", err)
			}
			defer newFile.Close()
			_, err = newFile.Write([]byte("[]"))
			if err != nil {
				return nil, fmt.Errorf("error initializing tasks.json file: %w", err)
			}
			return []model.Task{}, nil
		}

		return nil, fmt.Errorf("error opening tasks.json file: %w", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading the content of tasks.json file: %w", err)
	}

	if len(byteValue) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data (%s): %w", string(byteValue), err)
	}

	return tasks, nil
}

func WriteTaskToFile(tasks []model.Task) error {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("error opening tasks.json file: %v", err)
	}
	defer file.Close()

	file.Seek(0, 0)
	file.Truncate(0)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		return fmt.Errorf("error encoding tasks to file: %v", err)
	}
	return nil
}


func UpdateTask(taskID int, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		fmt.Printf("error reading tasks from file: %v", err)
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}
	return WriteTaskToFile(tasks)
}

func UpdateTaskStatus(taskID int, status string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		fmt.Printf("error reading tasks from file: %v", err)
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}
	return WriteTaskToFile(tasks)
}

func DeleteTask(taskID int) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		fmt.Printf("error reading tasks from file: %v", err)
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	return WriteTaskToFile(tasks)
}
