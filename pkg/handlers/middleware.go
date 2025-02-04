package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/cors"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, sessionKey)
		if err != nil {
			sendResponseWithError(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			sendResponseWithError(w, r, http.StatusForbidden, nil)
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
