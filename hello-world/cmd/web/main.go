package main

import (
	"fmt"
	"net/http"

	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//added to ignore browser favicon request that results in an extra cache hit
	http.HandleFunc("/favicon.ico", doNothing)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

func doNothing(w http.ResponseWriter, r *http.Request) {}