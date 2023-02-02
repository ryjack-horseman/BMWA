package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/config"
	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/favicon", http.HandlerFunc(doNothing))

	return mux
}
