package server

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

	userID, err := getUserSession(w, r)
	if err != nil {
		return
	}
	req := &dto.CreateTaskRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
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
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(task))
}

func (h *TaskHandler) getAll(w http.ResponseWriter, r *http.Request) {

	userID, err := getUserSession(w, r)
	if err != nil {
		return
	}

	tasks, err := h.TaskUseCase.GetAll(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(tasks))
}

func (h *TaskHandler) getByID(w http.ResponseWriter, r *http.Request) {
	//TO THINK: mayby I should to pass in argument userID
	taskID, err := getIDFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	task, err := h.TaskUseCase.GetByID(taskID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(task))
}

func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	taskID, err := getIDFromURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	if err := h.TaskUseCase.Delete(taskID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getUserSession(w http.ResponseWriter, r *http.Request) (int64, error) {
	session, err := sessionStore.Get(r, sessionKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return 0, err
	}
	userID := session.Values["user_id"].(int64)
	return userID, nil
}

func getIDFromURL(r *http.Request) (int64, error) {
	taskID := chi.URLParam(r, "taskID")
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return 0, err
	}
	return int64(taskIDInt), nil

}
