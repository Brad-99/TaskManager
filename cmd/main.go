// cmd/main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "TaskManager/internal/task"
)

// Handler for GET /tasks: List all tasks
func listTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task.GetTasks())
}

// Handler for POST /tasks: Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
    var newTask task.Task
    if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    createdTask := task.AddTask(newTask)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdTask)
}

// Handler for GET /tasks/{id}: Get a task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/tasks/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    t, found := task.GetTaskByID(id)
    if !found {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

// Handler for DELETE /tasks/{id}: Delete a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/tasks/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    if task.DeleteTask(id) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.NotFound(w, r)
    }
}

func main() {
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            listTasks(w, r)
        } else if r.Method == http.MethodPost {
            createTask(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getTask(w, r)
        } else if r.Method == http.MethodDelete {
            deleteTask(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server started on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
