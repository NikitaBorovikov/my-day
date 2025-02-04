package handlers

import (
	"net/http"
	"strconv"
	"toDoApp/pkg/dto"
	"toDoApp/pkg/model"
	"toDoApp/pkg/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TaskHandler struct {
	TaskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		TaskUseCase: taskUseCase,
	}
}

func (h *TaskHandler) registerRouters(r chi.Router) {
	r.Use(AuthMiddleware)

	r.Post("/", h.create)
	r.Get("/", h.getAll)
	r.Route("/{taskID}", func(r chi.Router) {
		r.Get("/", h.getByID)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

}

func (h *TaskHandler) create(w http.ResponseWriter, r *http.Request) {

	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}
	req := &dto.CreateTaskRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	task := &model.Task{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		IsImportant: req.IsImportant,
		DueDate:     req.DueDate,
	}

	if err := h.TaskUseCase.Create(task); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sendOKResponse(w, r, task)
}

func (h *TaskHandler) getAll(w http.ResponseWriter, r *http.Request) {

	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	tasks, err := h.TaskUseCase.GetAll(userID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	sendOKResponse(w, r, tasks)
}

func (h *TaskHandler) getByID(w http.ResponseWriter, r *http.Request) {
	taskID, err := getTaskIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.TaskUseCase.GetByID(taskID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	sendOKResponse(w, r, task)
}

func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request) {
	taskID, err := getTaskIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	req := &dto.UpdateTaskRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	task := &model.Task{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		IsImportant: req.IsImportant,
		IsDone:      req.IsDone,
		DueDate:     req.DueDate,
	}

	if err := h.TaskUseCase.Update(task); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sendOKResponse(w, r, task)
}

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	taskID, err := getTaskIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.TaskUseCase.Delete(taskID); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	sendOKResponse(w, r, nil)
}

func getTaskIDFromURL(r *http.Request) (int64, error) {
	taskID := chi.URLParam(r, "taskID")
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return 0, err
	}
	return int64(taskIDInt), nil
}
