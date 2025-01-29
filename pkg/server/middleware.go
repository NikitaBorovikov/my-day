package server

import (
	"context"
	"net/http"
	"toDoApp/pkg/dto"

	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, sessionKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, dto.NewResponse(err.Error()))
			return
		}

		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			w.WriteHeader(http.StatusForbidden)
			render.JSON(w, r, nil)
			return
		}

		userID := session.Values["user_id"].(int64)
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORSMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
}
