package handlers

import (
	"antibomberman/post/internal/adapters/rest/middlewares"
	"antibomberman/post/internal/di"
	"antibomberman/post/internal/models"
	"antibomberman/post/pkg/response"
	"antibomberman/post/pkg/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type postDi struct {
	*di.DI
}

func PostRouter(di *di.DI) chi.Router {
	d := &postDi{di}

	r := chi.NewRouter()
	r.Get("/", d.all)
	r.Get("/{id}", d.show)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTCheck(di.Cfg.JWTSecret))
		r.Post("/", d.create)
		r.Delete("/{id}", d.delete)

	})

	return r

}

func (d postDi) all(w http.ResponseWriter, r *http.Request) {
	posts, err := d.PostService.All()
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", posts)
}
func (d postDi) show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := d.PostService.GetById(id)
	if err != nil {

		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", post)
}
func (d postDi) create(w http.ResponseWriter, r *http.Request) {
	imagePath, err := storage.UploadImage(r)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	userId := r.Context().Value("user_id").(int)

	data := models.PostCreate{
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		ImagePath: imagePath,
		UserId:    userId,
	}
	post, err := d.PostService.Create(data)
	if err != nil {
		_ = storage.Delete(imagePath)

		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", post)
}
func (d postDi) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := d.PostService.Delete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", nil)
}
