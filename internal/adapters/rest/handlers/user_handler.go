package handlers

import (
	"antibomberman/post/internal/di"
	"antibomberman/post/internal/models"
	"antibomberman/post/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type userDi struct {
	*di.DI
}

func UserRouter(di *di.DI) chi.Router {
	u := &userDi{di}

	r := chi.NewRouter()
	r.Post("/register", u.register)
	r.Post("/login", u.login)
	return r

}

func (d *userDi) register(w http.ResponseWriter, r *http.Request) {

	validate := validator.New()

	data := models.UserCreate{
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
	}
	err := validate.Struct(data)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	err = d.UserService.Register(data)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", nil)
}
func (d *userDi) login(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	data := models.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	err := validate.Struct(data)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	token, err := d.UserService.Login(data)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", token)
}
