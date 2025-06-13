package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-service/internal/model"
	"task-service/internal/service"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// ListTasks godoc
// @Summary List all tasks
// @Description Get a list of all tasks, supports pagination and filtering by status
// @Tags tasks
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Param status query string false "Task status filter (Pending, InProgress, Completed)"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks [get]
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	page, size := parsePaginationParams(r)
	statusFilter := r.URL.Query().Get("status")

	tasks, totalCount, err := h.service.ListTasks(page, size, statusFilter)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tasks":       tasks,
		"total_count": totalCount,
		"page":        page,
		"size":        size,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Get details of a specific task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(uint(id))
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with title, description, status, and due date
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task to create"
// @Success 201 {object} model.Task
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to create task"
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update an existing task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body model.Task true "Updated task data"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid request payload or ID"
// @Failure 500 {string} string "Failed to update task"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask model.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateTask(uint(id), &updatedTask)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Failed to delete task"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteTask(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// --- Utility functions

func parsePaginationParams(r *http.Request) (int, int) {
	pageParam := r.URL.Query().Get("page")
	sizeParam := r.URL.Query().Get("size")

	page := 1
	size := 10

	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}

	if s, err := strconv.Atoi(sizeParam); err == nil && s > 0 {
		size = s
	}

	return page, size
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
