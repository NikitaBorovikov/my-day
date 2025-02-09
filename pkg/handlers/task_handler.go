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

// @Summary Create New Task
// @Security sessionKey
// @Tags tasks
// @Description create new task
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Task
// @Failure 400,401,403,422 {object} dto.Response
// @Param input body dto.CreateTaskRequest true "task info"
// @Router /task/ [post]
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

// @Summary Get All Tasks
// @Security sessionKey
// @Tags tasks
// @Description get all tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.Task
// @Failure 400,401,403 {object} dto.Response
// @Router /task/ [get]
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

// @Summary Get Task By ID
// @Security sessionKey
// @Tags tasks
// @Description get task by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Task
// @Failure 400,403 {object} dto.Response
// @Router /task/{taskID} [get]
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

// @Summary Update Task
// @Security sessionKey
// @Tags tasks
// @Description update task
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Task
// @Failure 400,403,422 {object} dto.Response
// @Param input body dto.UpdateTaskRequest true "task info"
// @Router /task/{taskID} [put]
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

// @Summary Delete Task
// @Security sessionKey
// @Tags tasks
// @Description delete task
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,403 {object} dto.Response
// @Router /task/{taskID} [delete]
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
