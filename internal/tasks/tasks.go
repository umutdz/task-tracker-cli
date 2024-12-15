package tasks

import (
	"fmt"
	"strconv"
	"task-tracker/internal/model"
	"task-tracker/internal/storage"
	"time"
)


func AddTask(description string) error {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading in AddTask")
	}
	var lastID int
	if len(tasks) > 0 {
		lastID = tasks[len(tasks)-1].ID
		lastID += 1
	} else {
		lastID = 1
	}
	newTask := model.Task{
		ID:          lastID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	return storage.WriteTaskToFile(tasks)
}


func UpdateTask(taskID string, description string) error {
	convertedID, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("error while converting")
	}
	return storage.UpdateTask(convertedID, description)
}

func DeleteTask(taskID string) error {
	convertedID, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("error while converting")
	}
	return storage.DeleteTask(convertedID)
}

func UpdateStatus(taskID string, newStatus string) error {
	validStatuses := map[string]string{
		"todo":        		"todo",
		"mark-in-progress": "in-progress",
		"mark-done":        "done",
	}
	if _, exist := validStatuses[newStatus]; !exist {
		return fmt.Errorf("invalid status: %s. Valid statuses are: mark-in-progress, mark-done", newStatus)
	}
	convertedID, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("error while converting")
	}
	return storage.UpdateTaskStatus(convertedID, validStatuses[newStatus])
}

func ListTasks(filter string) ([]model.Task, error) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading in AddTask")
	}
	if filter == "" {
		return tasks, nil
	}
	var filteredTasks []model.Task

	for _, task := range tasks {
		if task.Status == filter {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}
