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

// @Summary Create New Event
// @Security sessionKey
// @Tags events
// @Description create new event
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Event
// @Failure 400,401,403,422 {object} dto.Response
// @Param input body dto.CreateEventRequest true "event info"
// @Router /event/ [post]
func (h *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}
	req := &dto.CreateEventRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	event := &model.Event{
		UserID:        userID,
		Name:          req.Name,
		Description:   req.Description,
		AppointedDate: req.AppointedDate,
	}

	if err := h.EventUseCase.Create(event); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sendOKResponse(w, r, event)
}

// @Summary Get All Events
// @Security sessionKey
// @Tags events
// @Description get all events
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.Event
// @Failure 400,401,403 {object} dto.Response
// @Router /event/ [get]
func (h *EventHandler) getAll(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	events, err := h.EventUseCase.GetAll(userID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	sendOKResponse(w, r, events)
}

// @Summary Get Event By ID
// @Security sessionKey
// @Tags events
// @Description get event by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Event
// @Failure 400,403 {object} dto.Response
// @Router /event/{eventID} [get]
func (h *EventHandler) getByID(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	event, err := h.EventUseCase.GetByID(eventID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	sendOKResponse(w, r, event)
}

// @Summary Update Event
// @Security sessionKey
// @Tags events
// @Description update event
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Event
// @Failure 400,403,422 {object} dto.Response
// @Param input body dto.UpdateEventRequest true "event info"
// @Router /event/{eventID} [put]
func (h *EventHandler) update(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	req := &dto.UpdateEventRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	event := &model.Event{
		ID:            eventID,
		Name:          req.Name,
		Description:   req.Description,
		AppointedDate: req.AppointedDate,
	}

	if err := h.EventUseCase.Update(event); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sendOKResponse(w, r, event)
}

// @Summary Delete Event
// @Security sessionKey
// @Tags events
// @Description delete event
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,403 {object} dto.Response
// @Router /event/{eventID} [delete]
func (h *EventHandler) delete(w http.ResponseWriter, r *http.Request) {
	eventID, err := getEventIDFromURL(r)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.EventUseCase.Delete(eventID); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err.Error())
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
