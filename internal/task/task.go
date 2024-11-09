// internal/task/task.go
package task

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks = []Task{} // in-memory storage for tasks

// GetTasks returns the list of all tasks
func GetTasks() []Task {
	return tasks
}

// AddTask adds a new task to the list
func AddTask(newTask Task) Task {
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	return newTask
}

// GetTaskByID returns a task by ID
func GetTaskByID(id int) (Task, bool) {
	for _, t := range tasks {
		if t.ID == id {
			return t, true
		}
	}
	return Task{}, false
}

// DeleteTask removes a task by ID
func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
