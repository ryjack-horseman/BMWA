package handlers

import (
	"net/http"

	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/config"
	"github.com/ryjack-horseman/BMWA/hello-world/hello-world/pkg/render"
)

var Repo *Repository

//Repository is the repository type 
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
} 
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
// Home is the home page handler
func (m * Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func (m * Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
