package server

import (
	"net/http"
	"toDoApp/pkg/dto"
	"toDoApp/pkg/model"
	"toDoApp/pkg/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type EventHandler struct {
	EventUseCase *usecases.EventUseCase
}

func NewEventHandler(eventUseCase *usecases.EventUseCase) *EventHandler {
	return &EventHandler{
		EventUseCase: eventUseCase,
	}
}

func (h *EventHandler) registerRouters(r chi.Router) {
	r.Use(AuthMiddleware)
	r.Post("/", h.create)
	r.Get("/", h.getAll)

	r.Route("/{eventID}", func(r chi.Router) {
		r.Get("/", h.getByID)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

}

func (h *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, nil)
		return
	}
	req := &dto.CreateEventRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	event := &model.Event{
		UserID:        userID,
		Name:          req.Name,
		Description:   req.Description,
		AppointedDate: req.AppointedDate,
	}

	if err := h.EventUseCase.Create(event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(event))
}

func (h *EventHandler) getAll(w http.ResponseWriter, r *http.Request) {

}

func (h *EventHandler) getByID(w http.ResponseWriter, r *http.Request) {

}

func (h *EventHandler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *EventHandler) delete(w http.ResponseWriter, r *http.Request) {

}
