package rest

import (
	"antibomberman/post/internal/adapters/rest/handlers"
	"antibomberman/post/internal/di"
	"antibomberman/post/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type RouteHandler struct {
	di *di.DI
}

func New(di *di.DI) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.Error(w, "not found", http.StatusNotFound, nil)
	})
	r.MethodNotAllowed(func(writer http.ResponseWriter, request *http.Request) {
		response.Error(writer, "method not allowed", http.StatusMethodNotAllowed, nil)
	})

	r.Mount("/user", handlers.UserRouter(di))
	r.Mount("/post", handlers.PostRouter(di))
	r.Mount("/comment", handlers.CommentRouter(di))

	addr := ":" + di.Cfg.ServerPort
	log.Printf("server started at %s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return server
}
