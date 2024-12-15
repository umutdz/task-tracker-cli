package main

import (
	"fmt"
	"os"
	"strings"
	"task-tracker/internal/tasks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Usage: ./task-tracker <command>")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("No task provided. Usage: ./task-tracker add <task>")
			return
		}
		description := os.Args[2]
		tasks.AddTask(description)
		fmt.Printf("Adding task: %s\n", description)
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("No task ID provided. Usage: ./task-tracker update <task-id> <new-descriptionk>")
			return
		}
		taskID := os.Args[2]
		newDescription := strings.Join(os.Args[3:], " ")
		tasks.UpdateTask(taskID, newDescription)
		fmt.Printf("Updating task %s with new description: %s\n", taskID, newDescription)
	case "delete":
		if len(os.Args) < 2 {
			fmt.Println("No task ID provided. Usage: ./task-tracker delete <task-id>")
			return
		}
		taskID := os.Args[2]
		tasks.DeleteTask(taskID)
		fmt.Printf("Deleting task %s\n", taskID)
	case "mark-in-progress":
		if len(os.Args) < 2 {
			fmt.Println("No task ID provided. Usage: ./task-tracker mark-in-progress <task-id>")
			return
		}
		newStatus := os.Args[1]
		taskID := os.Args[2]
		tasks.UpdateStatus(taskID, newStatus)
		fmt.Printf("Marking task %s as in progress\n", taskID)
	case "mark-done":
		if len(os.Args) < 2 {
			fmt.Println("No task ID provided. Usage: ./task-tracker mark-done <task-id>")
			return
		}
		newStatus := os.Args[1]
		taskID := os.Args[2]
		tasks.UpdateStatus(taskID, newStatus)
		fmt.Printf("Marking task %s as done\n", taskID)
	case "list":
		if len(os.Args) == 2 {
			tasks, err := tasks.ListTasks("")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Listing all tasks:\n", tasks)
		} else if len(os.Args) == 3 {
			filter := os.Args[2]
			tasks, err := tasks.ListTasks(filter)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Listing tasks with filter:\n", filter, tasks)
		} else {
			fmt.Println("Invalid number of arguments. Usage: ./task-tracker list [filter]")
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
