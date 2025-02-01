package server

import (
	"net/http"
	"toDoApp/pkg/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type MyDayHandler struct {
	MyDayUseCase *usecases.MyDayUseCase
}

func NewMyDayHandler(myDayUseCase *usecases.MyDayUseCase) *MyDayHandler {
	return &MyDayHandler{
		MyDayUseCase: myDayUseCase,
	}
}

func (h *MyDayHandler) registerRouters(r chi.Router) {
	r.Use(AuthMiddleware)

	r.Get("/{date}", h.get)
}

func (h *MyDayHandler) get(w http.ResponseWriter, r *http.Request) {

	userID, ok := getUserID(r.Context())
	if !ok {
		logrus.Info("tut")
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	date := chi.URLParam(r, "date")

	myDay, err := h.MyDayUseCase.Get(userID, date)
	if err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	sendOKResponse(w, r, myDay)

}
