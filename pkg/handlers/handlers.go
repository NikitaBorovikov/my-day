package handlers

import (
	"context"
	"net/http"
	"toDoApp/pkg/dto"
	"toDoApp/pkg/usecases"

	_ "toDoApp/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handlers struct {
	userHandler  *UserHandler
	taskHandler  *TaskHandler
	eventHandler *EventHandler
	myDayHandler *MyDayHandler
}

func InitHandlers(u *usecases.UseCases) *Handlers {
	return &Handlers{
		userHandler:  NewUserHandler(u.UserUseCase),
		taskHandler:  NewTaskHandler(u.TaskUseCase),
		eventHandler: NewEventHandler(u.EventUseCase),
		myDayHandler: NewMyDayHandler(u.MyDayUseCase),
	}
}

func (h *Handlers) InitRouters() *chi.Mux {
	r := chi.NewRouter()
	r.Use(CORSMiddleware())

	r.Post("/reg", h.userHandler.signUp)
	r.Post("/login", h.userHandler.signIn)

	r.Route("/profile", func(r chi.Router) {
		h.userHandler.registerRouters(r)
	})

	r.Route("/task", func(r chi.Router) {
		h.taskHandler.registerRouters(r)
	})

	r.Route("/event", func(r chi.Router) {
		h.eventHandler.registerRouters(r)
	})

	r.Route("/myday", func(r chi.Router) {
		h.myDayHandler.registerRouters(r)
	})

	//swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}

func getUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

func sendResponseWithError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, dto.NewErrorResponse(err))
}

func sendOKResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewOKResponse(data))
}
