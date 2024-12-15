package tasks

import (
	"errors"
	"reflect"
)

func ValidateTaskDescription(description string) error {
	if description == "" {
		return errors.New("task description cannot be empty")
	}
	return nil
}

func ValidateTaskID(id int) error {
	IDType := reflect.TypeOf(id)
	if IDType.Kind() == reflect.Int {
		return errors.New("task ID must be a integer")
	}
	return nil
}

func ValidateTaskStatus(status string) error {
	validStatuses := []string{"todo", "in-progress", "done"}

	if status == "" {
		return errors.New("task status cannot be empty")
	}

	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}

	return errors.New("invalid task status: must be one of 'todo', 'in-progress', or 'done'")
}
