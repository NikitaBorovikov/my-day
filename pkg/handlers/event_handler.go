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
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}
	req := &dto.CreateEventRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	event := &model.Event{
		UserID:        userID,
		Name:          req.Name,
		Description:   req.Description,
		AppointedDate: req.AppointedDate,
	}

	if err := h.EventUseCase.Create(event); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	sendOKResponse(w, r, event)
}

func (h *EventHandler) getAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	events, err := h.EventUseCase.GetAll(userID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	sendOKResponse(w, r, events)
}

func (h *EventHandler) getByID(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	event, err := h.EventUseCase.GetByID(eventID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	sendOKResponse(w, r, event)
}

func (h *EventHandler) update(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	req := &dto.UpdateEventRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	event := &model.Event{
		ID:            eventID,
		Name:          req.Name,
		Description:   req.Description,
		AppointedDate: req.AppointedDate,
	}

	if err := h.EventUseCase.Update(event); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	sendOKResponse(w, r, event)
}

func (h *EventHandler) delete(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}
	if err := h.EventUseCase.Delete(eventID); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	sendOKResponse(w, r, nil)
}

func getEventIDFromURL(r *http.Request) (int64, error) {
	eventID := chi.URLParam(r, "eventID")
	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		return 0, err
	}

	return int64(eventIDInt), nil
}
