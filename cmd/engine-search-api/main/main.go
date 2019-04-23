package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yyh-gl/engine-search-api/api/router"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/", router.MainRouter())

	fmt.Println("========================")
	fmt.Println("Server Start @ http://localhost:3000")
	fmt.Println("========================")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
		return
	}
}
