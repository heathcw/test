package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Description string
	Done        bool
}

const taskFile = "tasks.json"

func main() {
	fmt.Println("Task Manager")
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Expected 'add', 'list', 'done', or 'delete")
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: todo add \"task description\"")
			return
		}
		err := addTask(args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		listTasks()
	case "done":
		if len(args) < 3 {
			fmt.Println("Usage: todo done <task number>")
			return
		}
		i, _ := strconv.Atoi(args[2])
		markDone(i)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: todo delete <task number>")
			return
		}
		i, _ := strconv.Atoi(args[2])
		deleteTask(i)
	default:
		fmt.Println("Unknown command:", args[1])
	}
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return tasks, nil
	}
	json.Unmarshal(file, &tasks)
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", " ")
	return os.WriteFile(taskFile, data, 0644)
}

func addTask(description string) error {
	tasks, _ := loadTasks()
	tasks = append(tasks, Task{Description: description})
	return saveTasks(tasks)
}

func listTasks() {
	tasks, _ := loadTasks()
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
}

func markDone(index int) error {
	tasks, _ := loadTasks()
	if index < 1 || index > len(tasks) {
		return fmt.Errorf("invalid task number")
	}
	tasks[index-1].Done = true
	return saveTasks(tasks)
}

func deleteTask(index int) error {
	tasks, _ := loadTasks()
	if index < 1 || index > len(tasks) {
		return fmt.Errorf("invalid task number")
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	return saveTasks(tasks)
}
