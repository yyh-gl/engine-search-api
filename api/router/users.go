package router

import (
	"github.com/yyh-gl/engine-search-api/api/handler/users"
	"net/http"

	"github.com/go-chi/chi"
)

func UsersRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", users.Index)
	return r
}
