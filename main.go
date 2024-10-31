package main

import (
	"ifood-backend-test/src/config"
	"ifood-backend-test/src/infra/api"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.LoadConfig()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	api.MakeSuggestionHandler(r)
	
	http.ListenAndServe(":9000", r)
}