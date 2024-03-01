package handlers

import (
	"antibomberman/post/internal/adapters/rest/middlewares"
	"antibomberman/post/internal/di"
	"github.com/go-chi/chi/v5"
	"net/http"
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

}
func (d commentDi) create(w http.ResponseWriter, r *http.Request) {

}
func (d commentDi) delete(w http.ResponseWriter, r *http.Request) {

}
