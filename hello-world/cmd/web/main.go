package main

import (
	"log"
	"net/http"

	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/config"
	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/handlers"
	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template Cache ")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func doNothing(w http.ResponseWriter, r *http.Request) {}
