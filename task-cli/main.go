package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Task represents a single task item
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

const tasksFile = "tasks.json"

// addTask adds a new task to the task list
// Demonstrates: pointer receiver to modify slice, auto-incrementing ID
func addTask(tasks *[]Task, description string) {
	// Find the next available ID
	maxID := 0
	for _, task := range *tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	newTask := Task{
		ID:          maxID + 1,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	*tasks = append(*tasks, newTask)
	fmt.Printf("Task added successfully! (ID: %d)\n", newTask.ID)
}

// listTasks displays all tasks in a readable format
// Demonstrates: range loop over slice, formatted output
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\n=== Task List ===")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[✓]"
		}
		fmt.Printf("%s ID: %d | %s | Created: %s\n",
			status,
			task.ID,
			task.Description,
			task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println()
}

// deleteTask removes a task by ID
// Demonstrates: error handling, slice manipulation
func deleteTask(tasks *[]Task, id int) error {
	for i, task := range *tasks {
		if task.ID == id {
			// Remove task by creating new slice without this element
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Printf("Task %d deleted successfully!\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// completeTask marks a task as completed
// Demonstrates: error handling, modifying struct in slice
func completeTask(tasks *[]Task, id int) error {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Completed = true
			fmt.Printf("Task %d marked as completed!\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// loadTasks loads tasks from the JSON file
// Demonstrates: file I/O, JSON unmarshalling, error handling
func loadTasks(filename string) ([]Task, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		// If file doesn't exist, return empty slice (not an error for first run)
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Unmarshal JSON data into Task slice
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return tasks, nil
}

// saveTasks saves tasks to the JSON file
// Demonstrates: JSON marshalling with indentation, file writing
func saveTasks(filename string, tasks []Task) error {
	// Marshal tasks to JSON with indentation for readability
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Write to file with appropriate permissions
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// showUsage displays the help message
func showUsage() {
	fmt.Println("\nTask Manager CLI - Go Learning Project")
	fmt.Println("\nUsage:")
	fmt.Println("  go run main.go <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add <description>   Add a new task")
	fmt.Println("  list                List all tasks")
	fmt.Println("  complete <id>       Mark a task as completed")
	fmt.Println("  delete <id>         Delete a task")
	fmt.Println("  help                Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go add \"Buy groceries\"")
	fmt.Println("  go run main.go list")
	fmt.Println("  go run main.go complete 1")
	fmt.Println("  go run main.go delete 1")
	fmt.Println()
}

func main() {
	// Check if command was provided
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	// Load existing tasks
	tasks, err := loadTasks(tasksFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	// Execute command
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			fmt.Println("Usage: go run main.go add <description>")
			os.Exit(1)
		}
		// Join all arguments after "add" as the description
		description := strings.Join(os.Args[2:], " ")
		addTask(&tasks, description)

		// Save tasks after adding
		if err := saveTasks(tasksFile, tasks); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}

	case "list":
		listTasks(tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: go run main.go complete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid task ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		if err := completeTask(&tasks, id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Save tasks after completing
		if err := saveTasks(tasksFile, tasks); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: go run main.go delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid task ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		if err := deleteTask(&tasks, id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Save tasks after deleting
		if err := saveTasks(tasksFile, tasks); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}

	case "help":
		showUsage()

	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		showUsage()
		os.Exit(1)
	}
}
