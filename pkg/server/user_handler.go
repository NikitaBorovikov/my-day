package server

import (
	"net/http"
	"toDoApp/pkg/dto"
	"toDoApp/pkg/model"
	"toDoApp/pkg/usecases"

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

func (h *UserHandler) signUp(w http.ResponseWriter, r *http.Request) {
	req := &dto.SignUpRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dto.NewResponse(err.Error())
		return
	}

	user := &model.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.UserUseCase.SignUp(user); err != nil {
		w.WriteHeader(http.StatusConflict)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(user))
}

func (h *UserHandler) signIn(w http.ResponseWriter, r *http.Request) {

	req := &dto.SignInRequest{}

	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	user, err := h.UserUseCase.SignIn(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	if err := setUserSession(w, r, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, dto.NewResponse(user.ID))

}

func logOut(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, sessionKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	cleanSessionInfo(session)

	if err := session.Save(r, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, dto.NewResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
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
		MaxAge:   3600, //seconds
		HttpOnly: true,
	}

	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}
