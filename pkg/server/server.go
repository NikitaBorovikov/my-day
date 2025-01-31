package server

import (
	"context"
	"net/http"
	"toDoApp/pkg/config"
	"toDoApp/pkg/dto"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
)

var (
	sessionKey   string
	sessionStore *sessions.CookieStore
)

type Handlers struct {
	userHandler  *UserHandler
	taskHandler  *TaskHandler
	eventHandler *EventHandler
}

func InitHandlers(userHandler *UserHandler, taskHandler *TaskHandler, eventHandler *EventHandler) *Handlers {
	return &Handlers{
		userHandler:  userHandler,
		taskHandler:  taskHandler,
		eventHandler: eventHandler,
	}
}

func Start(h *Handlers, cfg *config.Config) *chi.Mux {

	initSession(cfg.Http.SessionKey)

	server := initRouters(h)
	return server
}

func initRouters(h *Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(CORSMiddleware())
	r.Post("/reg", h.userHandler.signUp)
	r.Post("/login", h.userHandler.signIn)
	r.Post("/logout", logOut)

	r.Route("/task", func(r chi.Router) {
		h.taskHandler.registerRouters(r)
	})

	r.Route("/event", func(r chi.Router) {
		h.eventHandler.registerRouters(r)
	})

	return r
}

func initSession(key string) {
	sessionKey = key
	sessionStore = sessions.NewCookieStore([]byte(key))
}

func getUserID(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

func sendResponseWithError(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, dto.NewResponse(data))
}

func sendOKResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(data))
}
