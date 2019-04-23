package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func MainRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/users", UsersRouter())
	return r
}
