package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func UsersRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", hogeHundler)
	return r
}

func hogeHundler(w http.ResponseWriter, r *http.Request) {
	hoge := "test"
	w.Write([]byte(fmt.Sprintf("title:%s", hoge)))
}
