package server

import (
	"context"
	"net/http"
	"toDoApp/pkg/config"
	"toDoApp/pkg/handlers"

	"github.com/gorilla/sessions"
)

type Server struct {
	httpServer *http.Server
}

var (
	sessionKey   string
	sessionStore *sessions.CookieStore
)

func (s *Server) Run(h *handlers.Handlers, cfg *config.Config) error {

	handlers.InitSession(cfg.Http.SessionKey)

	router := h.InitRouters()

	s.httpServer = &http.Server{
		Addr:    cfg.Http.Port,
		Handler: router,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
