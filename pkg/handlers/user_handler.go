package handlers

import (
	"net/http"
	"toDoApp/pkg/dto"
	"toDoApp/pkg/model"
	"toDoApp/pkg/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (h *UserHandler) registerRouters(r chi.Router) {
	r.Use(AuthMiddleware)
	r.Get("/", h.get)
	r.Delete("/", h.delete)
	r.Post("/logout", logOut)
}

func (h *UserHandler) signUp(w http.ResponseWriter, r *http.Request) {
	req := &dto.SignUpRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	user := &model.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.UserUseCase.SignUp(user); err != nil {
		sendResponseWithError(w, r, http.StatusUnprocessableEntity, err)
		return
	}
	user.Password = ""
	sendOKResponse(w, r, user)
}

func (h *UserHandler) signIn(w http.ResponseWriter, r *http.Request) {

	req := &dto.SignInRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserUseCase.SignIn(req.Email, req.Password)
	if err != nil {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	if err := setUserSession(w, r, user); err != nil {
		sendResponseWithError(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, user.ID)

}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	userInfo, err := h.UserUseCase.Get(userID)
	if err != nil {
		sendResponseWithError(w, r, http.StatusBadRequest, err)
		return
	}

	sendOKResponse(w, r, userInfo)
}

func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r.Context())
	if !ok {
		sendResponseWithError(w, r, http.StatusUnauthorized, nil)
		return
	}

	go func() {
		if err := h.UserUseCase.Delete(userID); err != nil {
			sendResponseWithError(w, r, http.StatusBadRequest, err)
			return
		}
	}()

	go logOut(w, r)

	sendOKResponse(w, r, nil)
}

func logOut(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, sessionKey)
	if err != nil {
		sendResponseWithError(w, r, http.StatusInternalServerError, err)
		return
	}

	cleanSessionInfo(session)

	if err := session.Save(r, w); err != nil {
		sendResponseWithError(w, r, http.StatusInternalServerError, err)
		return
	}

	sendOKResponse(w, r, nil)
}

func cleanSessionInfo(session *sessions.Session) {
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1 // delete session
}

func setUserSession(w http.ResponseWriter, r *http.Request, user *model.User) error {
	session, err := sessionStore.Get(r, sessionKey)
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID

	session.Options = &sessions.Options{
		MaxAge:   3600 * 12, //seconds
		HttpOnly: true,
	}

	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}
