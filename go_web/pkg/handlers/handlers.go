package handlers

import (
	"net/http"

	"github.com/phanthehoang2503/go_training_ground/go_web/pkg/render"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
