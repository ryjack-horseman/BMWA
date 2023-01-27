package main

import (
	"fmt"
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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	//added to ignore browser favicon request that results in an extra cache hit
	http.HandleFunc("/favicon.ico", doNothing)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}
