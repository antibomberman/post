package handlers

import (
	"antibomberman/post/internal/adapters/rest/middlewares"
	"antibomberman/post/internal/di"
	"antibomberman/post/internal/models"
	"antibomberman/post/pkg/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type commentDi struct {
	*di.DI
}

func CommentRouter(di *di.DI) chi.Router {
	d := &commentDi{di}

	r := chi.NewRouter()
	r.Get("/", d.all)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTCheck(di.Cfg.JWTSecret))
		r.Post("/", d.create)
		r.Delete("/{id}", d.delete)
	})

	return r

}
func (d commentDi) all(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "post_id")
	comments, err := d.CommentService.GetByPostId(postId)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", comments)
}
func (d commentDi) create(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	data := models.CommentCreate{
		Content: r.FormValue("content"),
		PostID:  postId,
		UserID:  r.Context().Value("user_id").(int),
	}
	comment, err := d.CommentService.Create(data)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", comment)

}
func (d commentDi) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := d.CommentService.Delete(id)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	response.Success(w, "success", nil)
}
