package main

import (
	"github.com/Sirpyerre/go-course/pkg/config"
	"github.com/Sirpyerre/go-course/pkg/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
