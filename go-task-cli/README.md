# Go Task CLI

A command-line task manager application built in Go to demonstrate fundamental Go concepts like structs, slices, error handling, and JSON persistence - all without traditional OOP patterns.

## 🎯 Learning Objectives

This project demonstrates:
- **Structs** instead of classes for data modeling
- **Slices** for managing collections
- **Error handling** with Go's `if err != nil` pattern (no try-catch)
- **JSON marshalling/unmarshalling** for data persistence
- **Pointer receivers** for modifying data structures
- **File I/O** using the `os` package
- **CLI argument parsing** with `os.Args`

## 🚀 Features

- ✅ Add new tasks with descriptions
- ✅ List all tasks with status indicators
- ✅ Mark tasks as completed
- ✅ Delete tasks by ID
- ✅ Persistent storage using JSON file
- ✅ Auto-incrementing task IDs
- ✅ Timestamp tracking for task creation
- ✅ Clean error handling and user feedback

## 📁 Project Structure

```
go-task-cli/
├── main.go       # Complete application code
├── tasks.json    # Data persistence file (auto-created)
├── go.mod        # Go module file
└── README.md     # This file
```

## 🛠️ Installation

### Prerequisites
- Go 1.16 or higher

### Clone and Setup
```bash
cd go-task-cli
go build
```

## 📖 Usage

### Basic Commands

**Add a new task:**
```bash
go run main.go add "Buy groceries"
go run main.go add "Learn Go error handling"
```

**List all tasks:**
```bash
go run main.go list
```

Output example:
```
=== Task List ===
[ ] ID: 1 | Buy groceries | Created: 2026-01-31 16:27:18
[✓] ID: 2 | Learn Go error handling | Created: 2026-01-31 16:28:22
```

**Mark a task as completed:**
```bash
go run main.go complete 1
```

**Delete a task:**
```bash
go run main.go delete 1
```

**Show help:**
```bash
go run main.go help
```

### Using the Built Executable

After running `go build`, you can use the compiled binary:

**Windows:**
```bash
.\go-task-cli.exe add "My task"
.\go-task-cli.exe list
```

**Linux/Mac:**
```bash
./go-task-cli add "My task"
./go-task-cli list
```

## 🏗️ Code Structure

### Task Struct
```go
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
}
```

### Core Functions

- `addTask(*[]Task, string)` - Adds new task with auto-increment ID
- `listTasks([]Task)` - Displays all tasks in formatted output
- `deleteTask(*[]Task, int) error` - Removes task by ID
- `completeTask(*[]Task, int) error` - Marks task as completed
- `loadTasks(string) ([]Task, error)` - Loads tasks from JSON file
- `saveTasks(string, []Task) error` - Saves tasks to JSON file

## 📚 Go Concepts Demonstrated

### 1. Structs (Not Classes)
Go uses structs with methods instead of classes:
```go
type Task struct {
    ID          int
    Description string
    // ...
}
```

### 2. Error Handling Pattern
Go uses explicit error checking instead of exceptions:
```go
if err != nil {
    return fmt.Errorf("error message: %w", err)
}
```

### 3. Pointer Receivers
Functions that modify data use pointer receivers:
```go
func addTask(tasks *[]Task, description string) {
    *tasks = append(*tasks, newTask)
}
```

### 4. Slices for Collections
Go slices are dynamic arrays:
```go
var tasks []Task  // Not a class like List<Task>
tasks = append(tasks, newTask)
```

### 5. JSON Tags
Struct tags for JSON serialization:
```go
type Task struct {
    ID int `json:"id"`  // Maps to "id" in JSON
}
```

## 🔍 Key Differences from OOP

| Traditional OOP | Go Approach |
|----------------|-------------|
| Classes | Structs |
| Inheritance | Composition |
| Try-Catch | `if err != nil` |
| Exceptions | Error return values |
| Class methods | Functions with receivers |
| ArrayList/List | Slices |

## 📝 Data Persistence

Tasks are stored in `tasks.json` with human-readable formatting:

```json
[
  {
    "id": 1,
    "description": "Learn Go structs",
    "completed": false,
    "created_at": "2026-01-31T16:27:18+07:00"
  }
]
```

## 🧪 Testing

Manual testing steps:

```bash
# Test add
go run main.go add "Task 1"
go run main.go add "Task 2"

# Test list
go run main.go list

# Test complete
go run main.go complete 1

# Test delete
go run main.go delete 2

# Verify persistence (restart and list again)
go run main.go list
```

## 🎓 Learning Path

This project is perfect for:
1. **Go beginners** transitioning from OOP languages
2. Understanding **Go's approach to data structures**
3. Learning **Go's error handling philosophy**
4. Practicing **file I/O and JSON operations**
5. Building **practical CLI applications**

## 🤝 Contributing

This is a learning project. Feel free to:
- Extend functionality (priorities, due dates, categories)
- Add unit tests
- Implement additional commands
- Improve error messages
- Add data validation

## 📄 License

This project is created for educational purposes as part of learning Go.

## 🔗 Resources

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://go.dev/tour/)

---

**Built with ❤️ while learning Go**
